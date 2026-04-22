package repositories

import (
	"context"

	"github.com/Kayrit0/blog-api-go/internal/entities"
)

func (r *Repository) GetPosts() (*[]entities.Post, error) {
	posts := &[]entities.Post{}
	if err := r.db.QueryRow(context.Background(), "SELECT * FROM posts").Scan(posts); err != nil {
		return nil, err
	}
	return posts, nil
}
