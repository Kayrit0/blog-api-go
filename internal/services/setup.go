package services

import "github.com/Kayrit0/blog-api-go/internal/repositories"

type Service struct {
	repository *repositories.Repository
}

func Setup(repository *repositories.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
