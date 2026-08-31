package service

import (
	"errors"
	"fmt"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type OpeningService interface {
	List() ([]schemas.Opening, error)
	Show(id uint) (schemas.Opening, error)
	Create(opening *schemas.Opening) error
	Update(opening *schemas.Opening) error
	Delete(opening *schemas.Opening) error
	GetById(id string) (schemas.Opening, error)
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
	return s.repo.Create(opening)
}

func (s *openingService) Update(opening *schemas.Opening) error {
	return s.repo.Update(opening)
}

func (s *openingService) Delete(opening *schemas.Opening) error {
	return s.repo.Delete(opening)
}

func (s *openingService) GetById(id string) (schemas.Opening, error) {
	var openingID uint
	if _, err := fmt.Sscanf(id, "%d", &openingID); err != nil {
		return schemas.Opening{}, errors.New("invalid opening id")
	}
	return s.repo.Show(openingID)
}