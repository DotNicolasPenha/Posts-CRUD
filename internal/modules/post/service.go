package post

import (
	"errors"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
)

type Service struct {
	repository Repository
}

func NewService(r *Repository) *Service {
	if r == nil {
		logger.Fatal("repository of post service is nil")
	}
	return &Service{repository: *r}
}
func (s *Service) AddPost(post CreatePostDTO) error {
	if post.Username == "" {
		return errors.New("the username of post is empty")
	}
	if post.Body == "" {
		return errors.New("the body of post is empty")
	}
	if err := s.repository.Insert(post); err != nil {
		return err
	}
	return nil
}
func (s *Service) GetPosts() ([]Post, error) {
	posts, err := s.repository.FindMany()
	return posts, err
}
func (s *Service) GetOnePost(id string) (*Post, error) {
	if id == "" {
		return nil, errors.New("id param not found")
	}
	post, err := s.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) RemoveOne(id string) error {
	if id == "" {
		return errors.New("id param not found")
	}
	return s.repository.DeleteOne(id)
}
