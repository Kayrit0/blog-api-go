package services

import (
	"github.com/Kayrit0/blog-api-go/internal/entities"
)

func (s *Service) RegisterAccount(creds *entities.RegistrationCreds) error {
	// TODO: implement registration logic
	return nil
}

func (s *Service) LogInAccount(creds *entities.LogInCreds) (string, error) {
	// TODO: implement login logic
	return "", nil
}

func (s *Service) LogOutAccount(token string) error {
	// TODO: implement logout logic
	return nil
}
