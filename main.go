package main

import (
	"Image-loader/internal/repository"
	"Image-loader/internal/server"
	"Image-loader/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	var userRepo = repository.NewUserRepository("myNewFile.json")
	controller := service.NewController(userRepo)

	logger := logrus.New()

	srv := server.NewServer(":8000", logger, controller)
	srv.RegisterRoutes()

	srv.StartServer()
}
