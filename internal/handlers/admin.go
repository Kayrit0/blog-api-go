package handlers

import (
	"net/http"
	"strconv"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/Kayrit0/blog-api-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// UpdateUserRole changes the role of a user (owner only).
func (h *Handler) UpdateUserRole(ctx *gin.Context) {
	currentUser, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	userID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	// Prevent changing own role
	if currentUser.ID == uint(userID) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "cannot change your own role",
		})
		return
	}

	var req struct {
		Role entities.UserRole `json:"role" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.UpdateUserRole(uint(userID), req.Role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "role updated successfully",
	})
}
