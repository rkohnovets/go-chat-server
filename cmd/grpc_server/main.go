package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/rkohnovets/go-chat-server/api/chat_v1"
	serv "github.com/rkohnovets/go-chat-server/internal/grpc_server/chat_v1"
)

const grpcPort = 50051

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen at %d port: %v", grpcPort, err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &serv.Chat_v1_server{})

	log.Printf("grpc server starting to listen at %v", listener.Addr())

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
