package repository

import (
	"github.com/UxieGu1/gopportunities-api/schemas"
	"gorm.io/gorm"
)

type OpeningRepository struct {
	DB *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) *OpeningRepository {
	return &OpeningRepository{DB: db}
}

func (r *OpeningRepository) List() ([]schemas.Opening, error) {
	openings := []schemas.Opening{}
	if err := r.DB.Find(&openings).Error; err != nil {
		return nil, err
	}
	return openings, nil
}

func (r *OpeningRepository) Show(id uint) (schemas.Opening, error) {
	opening := schemas.Opening{}
	if err := r.DB.First(&opening, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return opening, nil
}

func (r *OpeningRepository) Create(opening *schemas.Opening) error {
	return r.DB.Create(opening).Error
}

func (r *OpeningRepository) Update(opening *schemas.Opening) error {
	return r.DB.Save(opening).Error
}

func (r *OpeningRepository) Delete(opening *schemas.Opening) error {
	return r.DB.Delete(opening).Error
}
