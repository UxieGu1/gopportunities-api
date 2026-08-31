package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(application *schemas.Application) error
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