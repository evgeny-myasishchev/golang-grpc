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

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("message: %v", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream: %v", info.FullMethod)
	return handler(srv, ss)
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptor),
		grpc.ChainStreamInterceptor(streamInterceptor),
	)
	chat.RegisterChatServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
