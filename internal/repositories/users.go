package repositories

import (
	"context"

	"github.com/Kayrit0/blog-api-go/internal/entities"
)

func (r *Repository) GetUsers(offset, limit int) ([]entities.User, error) {
	query := `SELECT id, username, email, password, role, created_at, updated_at FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUserByID(id uint) (*entities.User, error) {
	query := `SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE id = $1`
	var user entities.User
	err := r.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*entities.User, error) {
	query := `SELECT id, username, email, password, role, created_at, updated_at FROM users WHERE email = $1`
	var user entities.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *entities.User) error {
	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *Repository) UpdateUser(user *entities.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, updated_at = NOW() WHERE id = $4 RETURNING updated_at`
	return r.db.QueryRow(context.Background(), query, user.Username, user.Email, user.Password, user.ID).Scan(&user.UpdatedAt)
}

func (r *Repository) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

// UpdateUserRole updates only the role of a user.
func (r *Repository) UpdateUserRole(userID uint, role entities.UserRole) error {
	query := `UPDATE users SET role = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(context.Background(), query, role, userID)
	return err
}
