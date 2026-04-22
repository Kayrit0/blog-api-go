package main

import (
	"github.com/Kayrit0/blog-api-go/internal/database"
	"github.com/Kayrit0/blog-api-go/internal/handlers"
	"github.com/Kayrit0/blog-api-go/internal/libs"
	"github.com/Kayrit0/blog-api-go/internal/repositories"
	"github.com/Kayrit0/blog-api-go/internal/services"
)

func main() {
	cfg := libs.LoadConfig()

	dbPool := database.CreatePool(cfg)
	defer dbPool.Close()
	repo := repositories.Setup(dbPool)
	service := services.Setup(repo)
	routes := handlers.Setup(service)
	routes.Run(":8080")
}
