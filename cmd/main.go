package main

import (
	"github.com/Lanakod/go-todo"
	"github.com/Lanakod/go-todo/pkg/handler"
	"github.com/Lanakod/go-todo/pkg/repository"
	"github.com/Lanakod/go-todo/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured while initializing configs:\n%s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server:\n%s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
