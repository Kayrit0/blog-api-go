package handlers

import (
	"time"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/Kayrit0/blog-api-go/internal/middleware"
	"github.com/Kayrit0/blog-api-go/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

// Initialize your handlers here
func Setup(services *services.Service) *gin.Engine {
	h := &Handler{service: services}

	routes := gin.Default()

	// CORS configuration
	routes.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := routes.Group("/api/v1")
	{
		// Public auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", h.RegisterAccount)
			auth.POST("/login", h.LogInAccount)
		}

		// Authenticated routes
		authenticated := v1.Group("")
		authenticated.Use(middleware.RequireAuth())
		{
			authenticated.POST("/auth/logout", h.LogOutAccount)

			// Admin routes - owner only
			admin := authenticated.Group("/admin")
			admin.Use(middleware.RequireRole(entities.RoleOwner))
			{
				admin.PUT("/users/:id/role", h.UpdateUserRole)
			}

			// User routes - admin and owner only
			users := authenticated.Group("/users")
			users.Use(middleware.RequireRole(entities.RoleAdmin, entities.RoleOwner))
			{
				users.GET("/", h.GetAllUsers)
				users.GET("/:id", h.GetUserByID)
				users.PUT("/:id", h.UpdateUser)
				users.DELETE("/:id", h.DeleteUser)
			}

			// Post routes - authenticated users can create
			posts := authenticated.Group("/posts")
			{
				posts.POST("/", h.CreatePost)
				posts.PUT("/:id", h.UpdatePost)
				posts.DELETE("/:id", h.DeletePost)
			}
		}

		// Public post routes
		posts := v1.Group("/posts")
		{
			posts.GET("/", h.GetPosts)
			posts.GET("/:id", h.GetPostByID)
		}
	}

	routes.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	return routes
}
