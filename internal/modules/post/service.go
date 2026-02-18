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
func (s *Service) AddPost(post CreatePostDTO) []error {
	var errs []error
	if post.Username == "" {
		errs = append(errs, errors.New("the username of post is empty"))
	}
	if post.Body == "" {
		errs = append(errs, errors.New("the body of post is empty"))
	}
	if len(errs) > 0 {
		return errs
	}
	if err := s.repository.Insert(post); err != nil {
		return []error{err}
	}
	return nil
}
