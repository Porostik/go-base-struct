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

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "architecture/docs"
)

// @title Swagger API
// @version 1.0
// @description Swagger API for Golang Project.
// @termsOfService http://swagger.io/terms/

// @BasePath /api/

func main() {
	config := config.GetConfig()

	client, err := postgres.NewClient(context.TODO(), 5, config.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepositories(client)
	services := service.NewServices(repos)
	handlers := handler.NewHandler(services)

	router := handlers.Init()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server := http_server.NewServer(config, router)

	server.Start()
}
