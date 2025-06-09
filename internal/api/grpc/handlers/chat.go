package handlers

import (
	"context"
	"fmt"
	chat "github.com/ewik2k21/grpc_simple/pkg/proto/chat_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatHandler struct {
	chat.UnimplementedChatServiceServer
	//todo chat service
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *ChatHandler) CreateChat(ctx context.Context, req *chat.CreateChatRequest) (*chat.CreateChatResponse, error) {
	usernames := req.GetUsernames()
	if len(usernames) != 0 {
		return &chat.CreateChatResponse{Id: 1}, nil
	}

	return nil, fmt.Errorf("usernames = 0 ")
}

func (h *ChatHandler) DeleteChat(ctx context.Context, req *chat.DeleteChatRequest) (*emptypb.Empty, error) {
	id := req.GetId()
	if id != 0 {
		return nil, nil
	}
	return nil, fmt.Errorf("id = 0")
}

func (h *ChatHandler) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*emptypb.Empty, error) {
	text := req.GetText()
	if len(text) == 0 {
		return nil, fmt.Errorf("text is null")
	}

	return nil, nil
}
