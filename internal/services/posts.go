package services

import "github.com/Kayrit0/blog-api-go/internal/entities"

func (s *Service) GetAllPosts() ([]entities.Post, error) {
	return s.repository.GetPosts()
}

func (s *Service) GetPostByID(id uint) (*entities.Post, error) {
	return s.repository.GetPostByID(id)
}

func (s *Service) CreatePost(post *entities.Post) error {
	return s.repository.CreatePost(post)
}

func (s *Service) UpdatePost(post *entities.Post) error {
	return s.repository.UpdatePost(post)
}

func (s *Service) DeletePost(id uint) error {
	return s.repository.DeletePost(id)
}

func (s *Service) GetPostsByAuthorID(authorID uint) ([]entities.Post, error) {
	return s.repository.GetPostsByAuthorID(authorID)
}
