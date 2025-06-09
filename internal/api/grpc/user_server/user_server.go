package user_server

import (
	"github.com/ewik2k21/grpc_simple/internal/api/grpc/handlers"
	user "github.com/ewik2k21/grpc_simple/pkg/proto/user_api_v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{Port: port}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	userHandler := handlers.NewUserHandler()
	user.RegisterUserServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Errorf("server not started: %s", err)
		return err
	}
	return nil

}
