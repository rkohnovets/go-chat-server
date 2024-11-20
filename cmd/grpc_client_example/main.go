package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	utils "github.com/rkohnovets/go-auth/pkg/utils"
	desc "github.com/rkohnovets/go-chat-server/api/chat_v1"
)

func main() {
	grpc_server_address := "localhost:50051"
	conn, err := grpc.Dial(grpc_server_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &desc.ChatCreateRequest{
		Usernames: []string{"username1", "username2"},
	})
	if err != nil {
		log.Fatal("failed to create chat")
	}
	log.Printf("succesfully created chat, the new chat id: %d", r.GetId())

	result, err := utils.GetObjectJsonString(r)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	log.Println("full response object:\n" + string(result))
}
