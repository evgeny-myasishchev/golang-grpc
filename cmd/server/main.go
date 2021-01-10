package main

import (
	"log"
	"net"

	"github.com/evgeny-myasishchev/golang-grpc/pkg/chat"

	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatServiceServer
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