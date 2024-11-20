package chat_v1_server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"

	dest "github.com/rkohnovets/go-chat-server/api/chat_v1"
)

type Chat_v1_server struct {
	dest.UnimplementedChatV1Server
}

func (s *Chat_v1_server) Create(ctx context.Context, req *dest.ChatCreateRequest) (*dest.IdResponse, error) {
	chatId := gofakeit.Int64()
	log.Printf("creating chat, new id: %d", chatId)

	return &dest.IdResponse{
		Id: chatId,
	}, nil
}

func (s *Chat_v1_server) Delete(ctx context.Context, req *dest.IdRequest) (*emptypb.Empty, error) {
	chatId := req.GetId()
	log.Printf("deleting chat with id: %d", chatId)

	return &emptypb.Empty{}, nil
}

func (s *Chat_v1_server) SendMessage(ctx context.Context, req *dest.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("sending message (from %s): %s", req.GetFrom(), req.GetText())

	return &emptypb.Empty{}, nil
}
