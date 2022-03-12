package main

import (
	"architecture/internal/config"
	"architecture/internal/repository"
	"architecture/internal/service"
	http_server "architecture/internal/transport/http"
	"architecture/internal/transport/http/handler"
	"architecture/pkg/client/postgres"
	"context"
	"log"
)

func main() {
	config := config.GetConfig()

	client, err := postgres.NewClient(context.TODO(), 5, config.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepositories(client)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	server := http_server.NewServer(config, handlers.Init())

	server.Start()
}
