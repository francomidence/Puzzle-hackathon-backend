package main

import (
	"github.com/sirupsen/logrus"
	"puzzle-hackathon-backend/server/router"
)

func main() {

	logrus.Info("Starting the service...")

	server := router.NewRouter()
	serverError := server.Run()
	if serverError != nil {
		logrus.Fatalf("Server error: %v", serverError)
	}
}
