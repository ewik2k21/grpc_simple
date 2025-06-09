package handlers

import (
	"context"
	"fmt"
	user "github.com/ewik2k21/grpc_simple/pkg/proto/user_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer
	//todo chat service
}

func (u UserHandler) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	role := request.GetRole()
	if role.String() == "ADMIN" {
		return &user.CreateUserResponse{Id: 228}, nil
	}

	return nil, fmt.Errorf("user not admin")
}

func (u UserHandler) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	id := request.GetId()

	return &user.GetUserResponse{Id: id, Name: "star", Email: "asd@asd", Role: 1}, nil
}

func (u UserHandler) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	name := request.GetName()
	if name.String() == "" {
		return nil, fmt.Errorf("i want user name for id: %s ", id)
	}

	return nil, nil
}

func (u UserHandler) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*emptypb.Empty, error) {
	id := request.GetId()
	if id == 1 {
		return nil, nil
	}
	return nil, fmt.Errorf("id != 1")
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
