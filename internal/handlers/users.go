package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	user, err := h.service.GetUserByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	user, err := h.service.GetUserByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	var updateData struct {
		Username *string `json:"username" binding:"omitempty,min=3,max=50"`
		Email    *string `json:"email" binding:"omitempty,email"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if updateData.Username != nil {
		user.Username = *updateData.Username
	}
	if updateData.Email != nil {
		user.Email = *updateData.Email
	}

	if err := h.service.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	if err := h.service.DeleteUser(uint(userID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
