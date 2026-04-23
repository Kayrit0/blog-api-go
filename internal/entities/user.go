package entities

import "time"

// UserRole represents the role of a user in the system.
type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
	RoleOwner UserRole = "owner"
)

type User struct {
	ID        uint      `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      UserRole  `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Posts     []Post    `json:"posts,omitempty" db:"-"`
}
