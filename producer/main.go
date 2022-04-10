package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	protobuf "hw5/proto_files/proto"
)

type Message map[string]interface{}

func serialize(msg Message) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type server struct {
	protobuf.UnimplementedReturnResultFromWorkerServer
}

type result struct {
	len  int
	path []string
}

var q amqp.Queue
var ch *amqp.Channel

var id int
var mutex sync.Mutex
var idToChannel map[int]chan result

func (s *server) ReturnResult(ctx context.Context, in *protobuf.Result) (*protobuf.Empty, error) {
	fmt.Println("Got back result from receiver", in.Length, in.Path)

	// send result to channel
	mutex.Lock()
	channel := idToChannel[int(in.Id)]
	channel <- result{int(in.Length), in.Path}
	mutex.Unlock()

	return &protobuf.Empty{}, nil
}

func (s *server) RequestFromClient(ctx context.Context, in *protobuf.Request) (*protobuf.Result, error) {
	mutex.Lock()
	reqId := id
	id++

	channel := make(chan result, 1)
	idToChannel[reqId] = channel
	mutex.Unlock()

	fmt.Println("got request: ", reqId, in.From, in.To)

	// send message to rabbit mq
	msg := Message{"from": in.From, "to": in.To, "lang": in.Lang, "titles": in.Titles, "id": reqId}
	bytes, err := serialize(msg)
	failOnError(err, "failed to serialize message")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", bytes)

	// wait for result
	res := <-channel
	close(channel)

	fmt.Println("received ", res.len, res.path, "id: ", id)

	mutex.Lock()
	delete(idToChannel, reqId)
	mutex.Unlock()

	return &protobuf.Result{Length: int32(res.len), Path: res.path}, nil
}

var port = flag.Int("port", 50051, "The server port")
var rabbitMQAddr = flag.String("mqaddr", "amqp://guest:guest@localhost:5672/", "RabbitMQ address to connect to")

func main() {
	flag.Parse()

	idToChannel = make(map[int]chan result)

	// connect to rabbit mq
	var conn *amqp.Connection
	var err error
	for {
		conn, err = amqp.Dial(*rabbitMQAddr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err, "Failed to connect to RabbitMQ")

			time.Sleep(2 * time.Second)

			continue
		}
		defer conn.Close()
		break
	}

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err = ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// grpc stuff
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterReturnResultFromWorkerServer(s, &server{})

	// serve grpc server
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
