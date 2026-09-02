package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	List() ([]schemas.Opening, error)
	Show(id uint) (schemas.Opening, error)
	Create(opening *schemas.Opening) error
	CompanyBelongsToUser(companyID, userID uint) (bool, error)
	Update(opening *schemas.Opening) error
	Delete(opening *schemas.Opening) error
}

type openingRepository struct {
	db *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &openingRepository{db: db}
}

func (r *openingRepository) List() ([]schemas.Opening, error) {
	var openings []schemas.Opening
	if err := r.db.Find(&openings).Error; err != nil {
		return nil, err
	}
	return openings, nil
}

func (r *openingRepository) Show(id uint) (schemas.Opening, error) {
	var opening schemas.Opening
	if err := r.db.First(&opening, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return opening, nil
}

func (r *openingRepository) Create(opening *schemas.Opening) error {
	return r.db.Create(opening).Error
}

func (r *openingRepository) CompanyBelongsToUser(companyID, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&schemas.Company{}).
		Where("id = ? AND user_id = ?", companyID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *openingRepository) Update(opening *schemas.Opening) error {
	return r.db.Save(opening).Error
}

func (r *openingRepository) Delete(opening *schemas.Opening) error {
	return r.db.Delete(opening).Error
}
