package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protobuf "hw5/proto_files/proto"
	receive "hw5/worker/receive"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// sender addr
var addr = flag.String("addr", "localhost:50051", "sender address to return result to")
var rabbitMQAddr = flag.String("mqaddr", "amqp://guest:guest@localhost:5672/", "RabbitMQ address to connect to")

func main() {
	flag.Parse()

	// grpc stuff to return result back to sender

	// Set up a connection to the server.
	conn1, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c := protobuf.NewReturnResultFromWorkerClient(conn1)

	// connect to rabbit mq
	var conn *amqp.Connection
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

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		// receive messages
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msg, err := receive.Deserialize(d.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "deserialization error", msg, err)
				continue
			}
			from := msg["from"].(string)
			to := msg["to"].(string)
			lang := msg["lang"].(string)
			titles := msg["titles"].(bool)
			id := msg["id"].(float64)

			length, path := receive.FindPath(from, to, lang, titles)
			fmt.Println(length, path)

			// return result to sender
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			_, err = c.ReturnResult(ctx, &protobuf.Result{Length: int32(length), Path: path, Id: int32(id)})

			if err != nil {
				fmt.Fprintln(os.Stderr, "grpc error", err)
				continue
			}
			// fmt.Println(extractLinks(from))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
