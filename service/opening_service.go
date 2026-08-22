package service

import (
	"errors"
	"fmt"

	"github.com/UxieGu1/gopportunities-api/repository"
	"github.com/UxieGu1/gopportunities-api/schemas"
	"gorm.io/gorm"
)

type OpeningService struct {
	Repository *repository.OpeningRepository
}

func NewOpeningService(db *gorm.DB) *OpeningService {
	return &OpeningService{Repository: repository.NewOpeningRepository(db)}
}

func (s *OpeningService) List() ([]schemas.Opening, error) {
	return s.Repository.List()
}

func (s *OpeningService) Show(id uint) (schemas.Opening, error) {
	return s.Repository.Show(id)
}

func (s *OpeningService) Create(opening *schemas.Opening) error {
	return s.Repository.Create(opening)
}

func (s *OpeningService) Update(opening *schemas.Opening) error {
	return s.Repository.Update(opening)
}

func (s *OpeningService) Delete(opening *schemas.Opening) error {
	return s.Repository.Delete(opening)
}

func (s *OpeningService) GetByID(id string) (schemas.Opening, error) {
	var openingID uint
	if _, err := fmt.Sscanf(id, "%d", &openingID); err != nil {
		return schemas.Opening{}, errors.New("invalid opening id")
	}
	return s.Repository.Show(openingID)
}
