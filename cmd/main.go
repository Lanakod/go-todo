package main

import (
	"github.com/Lanakod/go-todo"
	"github.com/Lanakod/go-todo/pkg/handler"
	"github.com/Lanakod/go-todo/pkg/repository"
	"github.com/Lanakod/go-todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server:\n%s", err.Error())
	}
}
