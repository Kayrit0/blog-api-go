package services

import (
	"errors"

	"github.com/Kayrit0/blog-api-go/internal/entities"
)

// UpdateUserRole updates the role of a user.
func (s *Service) UpdateUserRole(userID uint, newRole entities.UserRole) error {
	// Validate role
	if newRole != entities.RoleUser && newRole != entities.RoleAdmin && newRole != entities.RoleOwner {
		return errors.New("invalid role")
	}

	return s.repository.UpdateUserRole(userID, newRole)
}
