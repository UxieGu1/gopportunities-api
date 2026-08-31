package service

import (
	"errors"
	"fmt"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type UserService interface {
	Create(user *schemas.User) error
	GetByID(id uint) (*schemas.User, error)
	GetByEmail(email string) (*schemas.User, error)
	GetAll() ([]schemas.User, error)
	Update(user *schemas.User) error
	Delete(id uint) error
	GetByIdStr(id string) (*schemas.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *schemas.User) error {
	if user.Email == "" {
		return errors.New("o e-mail é obrigatório")
	}
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*schemas.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) GetByEmail(email string) (*schemas.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) GetAll() ([]schemas.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Update(user *schemas.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) GetByIdStr(id string) (*schemas.User, error) {
	var userID uint
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		return nil, errors.New("invalid user id format")
	}
	return s.repo.GetById(userID)
}