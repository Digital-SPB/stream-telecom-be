package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handlers "github.com/greenblat17/stream-telecom/internal/handler"
	"github.com/greenblat17/stream-telecom/internal/repo"
	"github.com/greenblat17/stream-telecom/internal/service"
	"github.com/greenblat17/stream-telecom/pkg/httpserver"
	"github.com/sirupsen/logrus"
)

func main() {
	Run()
}

func Run() {
	//logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Repositories
	repos := repo.NewRepository()

	// Service
	services := service.NewService(repos)

	// Handlers
	handlers := handlers.NewHandler(services)
	// handlers := handle/r.NewHandler(services)
	//handlers := handler.NewHandler(services)

	//HTTP server

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()

	logrus.Print("hui zalupa i pizda")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error while server shutting down: %s", err.Error())
	}

}
