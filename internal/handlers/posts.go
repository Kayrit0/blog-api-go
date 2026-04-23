package handlers

import (
	"net/http"
	"strconv"

	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/Kayrit0/blog-api-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(ctx *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPostByID(ctx *gin.Context) {
	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post ID",
		})
		return
	}

	post, err := h.service.GetPostByID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePost(ctx *gin.Context) {
	currentUser, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	var req struct {
		Title   string `json:"title" binding:"required,min=3,max=255"`
		Content string `json:"content" binding:"required,min=10"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := &entities.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: currentUser.ID,
	}

	if err := h.service.CreatePost(post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, post)
}

func (h *Handler) UpdatePost(ctx *gin.Context) {
	currentUser, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post ID",
		})
		return
	}

	post, err := h.service.GetPostByID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	}

	// Check authorization: author  OR owner
	if post.AuthorID != currentUser.ID &&
		currentUser.Role != entities.RoleOwner {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
		return
	}

	var req struct {
		Title   *string `json:"title" binding:"omitempty,min=3,max=255"`
		Content *string `json:"content" binding:"omitempty,min=10"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update only provided fields
	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Content != nil {
		post.Content = *req.Content
	}

	if err := h.service.UpdatePost(post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(ctx *gin.Context) {
	currentUser, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	postID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post ID",
		})
		return
	}

	post, err := h.service.GetPostByID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	}

	// Check authorization: author OR admin OR owner
	if post.AuthorID != currentUser.ID &&
		currentUser.Role != entities.RoleAdmin &&
		currentUser.Role != entities.RoleOwner {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
		return
	}

	if err := h.service.DeletePost(uint(postID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post deleted successfully",
	})
}
