package handlers

import (
	"github.com/Kayrit0/blog-api-go/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

// Initialize your handlers here
func Setup(services *services.Service) *gin.Engine {
	h := &Handler{service: services}

	routes := gin.Default()

	v1 := routes.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", h.RegisterAccount)
			auth.POST("/login", h.LogInAccount)
			auth.POST("/logout", h.LogOutAccount)
		}

		users := v1.Group("/users")
		{
			users.GET("/", h.GetAllUsers)
			users.GET("/:id")
			users.PUT("/:id")
		}

		posts := v1.Group("/posts")
		{
			posts.GET("/", h.GetPosts)
			posts.POST("/")
			posts.GET("/:id")
			posts.PUT("/:id")
			posts.DELETE("/:id")
		}
	}

	routes.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	return routes
}
