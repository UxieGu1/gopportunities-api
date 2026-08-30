package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company *schemas.Company) error
	GetById(id uint) (*schemas.Company, error)
	GetAll() ([]schemas.Company, error)
	Update(company *schemas.Company) error
	Delete(id uint) error
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(company *schemas.Company) error {
	return r.db.Create(company).Error
}

func (r *companyRepository) GetById(id uint) (*schemas.Company, error) {
	var company schemas.Company

	err := r.db.Preload("Openings").First(&company, id).Error
	if err != nil{
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) GetAll() ([]schemas.Company, error) {
	var companies []schemas.Company

	err := r.db.Preload("Openings").Find(&companies).Error
	if err != nil{
		return nil, err
	}
	return companies, nil
}

func (r *companyRepository) Update(company *schemas.Company) error {
	return r.db.Save(company).Error
}

func (r *companyRepository) Delete(id uint) error {
	return r.db.Delete(&schemas.Company{}, id).Error
}





