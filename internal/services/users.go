package services

import "github.com/Kayrit0/blog-api-go/internal/entities"

func (s *Service) GetAllUsers() ([]entities.User, error) {
	return s.repository.GetUsers(0, 50)
}

func (s *Service) GetUserByID(id uint) (*entities.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *Service) UpdateUser(updatedData *entities.User) error {
	return s.repository.UpdateUser(updatedData)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repository.DeleteUser(id)
}
