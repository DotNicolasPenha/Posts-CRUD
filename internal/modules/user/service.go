package post

import (
	"errors"

	"github.com/DotNicolasPenha/Posts-CRUD/internal/common/logger"
)

type Service struct {
	repository Repository
}

func NewService(repository *Repository) *Service {
	if repository == nil {
		logger.Fatal("repository of user service is nil")
	}
	return &Service{
		repository: *repository,
	}
}

func (s *Service) AddUser(createUserDto CreateUserDTO) error {
	if createUserDto.Username == "" {
		return errors.New("username is undefined")
	}
	if createUserDto.Bio == "" {
		return errors.New("bio is undefined")
	}
	return s.repository.Insert(createUserDto)
}
