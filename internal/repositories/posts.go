package repositories

import (
	"context"

	"github.com/Kayrit0/blog-api-go/internal/entities"
)

func (r *Repository) GetPosts() ([]entities.Post, error) {
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM posts ORDER BY created_at DESC`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []entities.Post{}
	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Repository) GetPostByID(id uint) (*entities.Post, error) {
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE id = $1`
	var post entities.Post
	err := r.db.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Repository) CreatePost(post *entities.Post) error {
	query := `INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(context.Background(), query, post.Title, post.Content, post.AuthorID).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}

func (r *Repository) UpdatePost(post *entities.Post) error {
	query := `UPDATE posts SET title = $1, content = $2, updated_at = NOW() WHERE id = $3 RETURNING updated_at`
	return r.db.QueryRow(context.Background(), query, post.Title, post.Content, post.ID).Scan(&post.UpdatedAt)
}

func (r *Repository) DeletePost(id uint) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *Repository) GetPostsByAuthorID(authorID uint) ([]entities.Post, error) {
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE author_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.Query(context.Background(), query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []entities.Post{}
	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
