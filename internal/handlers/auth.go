package handlers

import (
	"net/http"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterAccount(ctx *gin.Context) {
	if token, err := ctx.Cookie("token"); err == nil && token != "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "already logged in"})
		return
	}

	creds := &entities.RegistrationCreds{}

	if err := ctx.ShouldBindJSON(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.RegisterAccount(creds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", token, 86400, "/", "", false, true)
	ctx.JSON(http.StatusCreated, gin.H{"message": "account created successfully"})
}

func (h *Handler) LogInAccount(ctx *gin.Context) {
	if token, err := ctx.Cookie("token"); err == nil && token != "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "already logged in"})
		return
	}

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

	ctx.SetCookie("token", token, 86400, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logged in successfully"})
}

func (h *Handler) LogOutAccount(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
