package main

import (
	"github.com/Kayrit0/blog-api-go/internal/handlers"
	"github.com/Kayrit0/blog-api-go/internal/repositories"
	"github.com/Kayrit0/blog-api-go/internal/services"
)

func main() {

	repo := repositories.Setup()
	service := services.Setup(repo)
	routes := handlers.Setup(service)
	routes.Run(":8080")
}
