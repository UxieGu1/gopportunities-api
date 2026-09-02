package service

import (
	"errors"
	"fmt"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user *schemas.User, rawPassword string) error
	GetByID(id uint) (*schemas.User, error)
	GetByEmail(email string) (*schemas.User, error)
	GetAll() ([]schemas.User, error)
	Update(user *schemas.User, rawPassword string) error
	Delete(id uint) error
	GetByIdStr(id string) (*schemas.User, error)
	Login(email, password string) (*schemas.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *schemas.User, rawPassword string) error {
	if user.Email == "" {
		return errors.New("o e-mail é obrigatório")
	}
	if rawPassword == "" {
		return errors.New("a senha é obrigatória")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return fmt.Errorf("falha ao criptografar a senha: %v", err)
	}
	user.PasswordHash = string(hashedBytes)
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*schemas.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetByEmail(email string) (*schemas.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) GetAll() ([]schemas.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Update(user *schemas.User, rawPassword string) error {
	if rawPassword != "" {
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
		if err != nil {
			return fmt.Errorf("falha ao criptografar a senha: %v", err)
		}
		user.PasswordHash = string(hashedBytes)
	}
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
	return s.repo.GetByID(userID)
}

func (s *userService) Login(email, rawPassword string) (*schemas.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(rawPassword))
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	return user, nil
}
