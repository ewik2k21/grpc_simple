package config

import "github.com/joho/godotenv"
import "github.com/sirupsen/logrus"

func LoadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Fatalf("error loading .env file")
	}

}
