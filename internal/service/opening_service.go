package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type OpeningService interface {
	List() ([]schemas.Opening, error)
	Show(id uint) (schemas.Opening, error)
	Create(opening *schemas.Opening) error
	CreateForUser(opening *schemas.Opening, userID uint) error
	Update(opening *schemas.Opening) error
	UpdateForUser(opening *schemas.Opening, userID uint) error
	Delete(opening *schemas.Opening) error
	DeleteForUser(opening *schemas.Opening, userID uint) error
	GetByID(id string) (schemas.Opening, error)
}

type openingService struct {
	repo repository.OpeningRepository
}

func NewOpeningService(repo repository.OpeningRepository) OpeningService {
	return &openingService{repo: repo}
}

func (s *openingService) List() ([]schemas.Opening, error) {
	return s.repo.List()
}

func (s *openingService) Show(id uint) (schemas.Opening, error) {
	return s.repo.Show(id)
}

func (s *openingService) Create(opening *schemas.Opening) error {
	if opening.Status == "" {
		opening.Status = "OPEN"
	}
	opening.Status = strings.ToUpper(strings.TrimSpace(opening.Status))
	if opening.Status != "OPEN" {
		return errors.New("uma nova vaga deve iniciar com status OPEN")
	}
	return s.repo.Create(opening)
}

func (s *openingService) CreateForUser(opening *schemas.Opening, userID uint) error {
	belongs, err := s.repo.CompanyBelongsToUser(opening.CompanyID, userID)
	if err != nil {
		return err
	}
	if !belongs {
		return errors.New("empresa não pertence ao usuário autenticado")
	}
	return s.repo.Create(opening)
}

func (s *openingService) Update(opening *schemas.Opening) error {
	if err := validateOpeningStatus(opening.Status); err != nil {
		return err
	}
	return s.repo.Update(opening)
}

func (s *openingService) UpdateForUser(opening *schemas.Opening, userID uint) error {
	if err := validateOpeningStatus(opening.Status); err != nil {
		return err
	}
	belongs, err := s.repo.CompanyBelongsToUser(opening.CompanyID, userID)
	if err != nil {
		return err
	}
	if !belongs {
		return errors.New("vaga não pertence ao usuário autenticado")
	}
	return s.repo.Update(opening)
}

func validateOpeningStatus(status string) error {
	status = strings.ToUpper(strings.TrimSpace(status))
	if status != "OPEN" && status != "PAUSED" && status != "CLOSED" {
		return errors.New("status de vaga inválido")
	}
	return nil
}

func (s *openingService) Delete(opening *schemas.Opening) error {
	return s.repo.Delete(opening)
}

func (s *openingService) DeleteForUser(opening *schemas.Opening, userID uint) error {
	belongs, err := s.repo.CompanyBelongsToUser(opening.CompanyID, userID)
	if err != nil {
		return err
	}
	if !belongs {
		return errors.New("vaga não pertence ao usuário autenticado")
	}
	return s.repo.Delete(opening)
}

func (s *openingService) GetByID(id string) (schemas.Opening, error) {
	var openingID uint
	if _, err := fmt.Sscanf(id, "%d", &openingID); err != nil {
		return schemas.Opening{}, errors.New("invalid opening id")
	}
	return s.repo.Show(openingID)
}
