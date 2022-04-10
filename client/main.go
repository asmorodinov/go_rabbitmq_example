package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protobuf "hw5/proto_files/proto"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	from := os.Args[1]
	to := os.Args[2]
	lang := os.Args[3]
	titles := os.Args[4] == "true"
	addr := os.Args[5]
	if addr == "" {
		addr = "localhost:50051"
	}

	// grpc stuff to return result back to sender

	// Set up a connection to the server.
	fmt.Println("addr:", addr)
	conn1, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c := protobuf.NewReturnResultFromWorkerClient(conn1)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("Waiting for result...")
	res, err := c.RequestFromClient(ctx, &protobuf.Request{From: from, To: to, Lang: lang, Titles: titles})

	if err != nil {
		fmt.Fprintln(os.Stderr, "grpc error", err)
		return
	}

	fmt.Println("Got result:", res.Length, res.Path)
}
