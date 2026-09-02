package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(application *schemas.Application) error
	Exists(candidateID, openingID uint) (bool, error)
	CandidateExists(candidateID uint) (bool, error)
	OpeningExists(openingID uint) (bool, error)
	GetOpeningStatus(openingID uint) (string, error)
	GetByID(id uint) (*schemas.Application, error)
	GetAll() ([]schemas.Application, error)
	Update(application *schemas.Application) error
	Delete(id uint) error
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db: db}
}

func (r *applicationRepository) Create(application *schemas.Application) error {
	return r.db.Create(application).Error
}

func (r *applicationRepository) Exists(candidateID, openingID uint) (bool, error) {
	var count int64
	err := r.db.Model(&schemas.Application{}).
		Where("candidate_id = ? AND opening_id = ?", candidateID, openingID).
		Count(&count).Error
	return count > 0, err
}

func (r *applicationRepository) CandidateExists(candidateID uint) (bool, error) {
	var count int64
	err := r.db.Model(&schemas.Candidate{}).Where("id = ?", candidateID).Count(&count).Error
	return count > 0, err
}

func (r *applicationRepository) OpeningExists(openingID uint) (bool, error) {
	var count int64
	err := r.db.Model(&schemas.Opening{}).Where("id = ?", openingID).Count(&count).Error
	return count > 0, err
}

func (r *applicationRepository) GetOpeningStatus(openingID uint) (string, error) {
	var opening schemas.Opening
	if err := r.db.Select("status").First(&opening, openingID).Error; err != nil {
		return "", err
	}
	return opening.Status, nil
}

func (r *applicationRepository) GetByID(id uint) (*schemas.Application, error) {
	var application schemas.Application
	if err := r.db.Preload("Candidate").Preload("Opening").First(&application, id).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *applicationRepository) GetAll() ([]schemas.Application, error) {
	var applications []schemas.Application
	if err := r.db.Preload("Candidate").Preload("Opening").Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (r *applicationRepository) Update(application *schemas.Application) error {
	return r.db.Save(application).Error
}

func (r *applicationRepository) Delete(id uint) error {
	return r.db.Delete(&schemas.Application{}, id).Error
}
