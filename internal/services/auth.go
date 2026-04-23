package services

import (
	"github.com/Kayrit0/blog-api-go/internal/entities"
	"github.com/Kayrit0/blog-api-go/internal/libs"
)

func (s *Service) RegisterAccount(creds *entities.RegistrationCreds) (string, error) {
	hashedPass, err := libs.HashPass(creds.Password)
	if err != nil {
		return "", err
	}

	user := &entities.User{
		Email:    creds.Email,
		Username: creds.Username,
		Password: hashedPass,
		Role:     entities.RoleUser,
	}

	if err := s.repository.CreateUser(user); err != nil {
		return "", err
	}

	token, err := libs.CreateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) LogInAccount(creds *entities.LogInCreds) (string, error) {
	user, err := s.repository.GetUserByEmail(creds.Email)
	if err != nil {
		return "", err
	}

	if err := libs.ComparePass(creds.Password, user.Password); err != nil {
		return "", err
	}

	token, err := libs.CreateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
