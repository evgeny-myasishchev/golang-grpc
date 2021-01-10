package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/evgeny-myasishchev/golang-grpc/pkg/chat"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &chat.Message{Body: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Body)

	log.Println("Getting messages")

	stream, err := c.GetMessages(ctx, &chat.GetMessagesRequest{})
	if err != nil {
		log.Fatalf("could not get messages: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("Done")
			break
		}
		log.Println("Got message:", msg.Body)
	}
}
