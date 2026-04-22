package services

import "github.com/Kayrit0/blog-api-go/internal/entities"

func (s *Service) GetAllUsers() ([]entities.User, error) {
	return s.repository.GetUsers(0, 50)
}
