package handlers

import (
	"net/http"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterAccount(ctx *gin.Context) {
	creds := &entities.RegistrationCreds{}

	if err := ctx.ShouldBindJSON(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RegisterAccount(creds); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "account created successfully"})
}

func (h *Handler) LogInAccount(ctx *gin.Context) {
	creds := &entities.LogInCreds{}

	if err := ctx.ShouldBindJSON(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.LogInAccount(creds)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) LogOutAccount(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
		return
	}

	if err := h.service.LogOutAccount(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
