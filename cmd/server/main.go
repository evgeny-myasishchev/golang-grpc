package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/evgeny-myasishchev/golang-grpc/pkg/chat"

	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatServiceServer
}

func (s *server) SayHello(_ context.Context, in *chat.Message) (*chat.Message, error) {
	return &chat.Message{Body: in.Body + "World"}, nil
}

func (s *server) GetMessages(_ *chat.GetMessagesRequest, out chat.ChatService_GetMessagesServer) error {
	for i := 0; i < 10; i++ {
		if err := out.Send(&chat.Message{Body: fmt.Sprintf("Message #%v", i)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
