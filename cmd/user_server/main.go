package main

import (
	"github.com/ewik2k21/grpc_simple/config"
	"github.com/ewik2k21/grpc_simple/internal/api/grpc/user_server"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	config.LoadEnv()
	server := user_server.NewServer(os.Getenv(config.GRPCPort2))

	if err := server.Run(); err != nil {
		logrus.Fatalf("failed run server: %s", err)
	}
}
