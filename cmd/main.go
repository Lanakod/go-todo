package main

import (
	"context"
	"github.com/Lanakod/go-todo"
	"github.com/Lanakod/go-todo/pkg/handler"
	"github.com/Lanakod/go-todo/pkg/repository"
	"github.com/Lanakod/go-todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occured while initializing configs:\n%s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error occured while loading env variables:\n%s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Error occured while initializing database:\n%s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server:\n%s", err.Error())
		}
	}()

	logrus.Printf("TodoApp started on port - %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Print("TodoApp Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down:\n%s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on closing db connection:\n%s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
