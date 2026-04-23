package middleware

import (
	"errors"
	"net/http"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/Kayrit0/blog-api-go/internal/libs"
	"github.com/gin-gonic/gin"
)

// RequireAuth validates JWT token and extracts user info to context.
func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			ctx.Abort()
			return
		}

		user, err := libs.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

// RequireRole checks if user has one of the required roles.
func RequireRole(roles ...entities.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := GetUserFromContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if user.Role == role {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
		ctx.Abort()
	}
}

// GetUserFromContext extracts user from gin context.
func GetUserFromContext(ctx *gin.Context) (*entities.User, error) {
	value, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("user not found in context")
	}

	user, ok := value.(entities.User)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return &user, nil
}
