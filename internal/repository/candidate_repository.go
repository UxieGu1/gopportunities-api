package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)

type CandidateRepository interface {
	Create(candidate *schemas.Candidate) error
	GetByID(id uint) (*schemas.Candidate, error)
	GetByIDForUser(id, userID uint) (*schemas.Candidate, error)
	GetByUserID(userID uint) (*schemas.Candidate, error)
	GetByEmail(email string) (*schemas.Candidate, error)
	GetAll() ([]schemas.Candidate, error)
	Update(candidate *schemas.Candidate) error
	Delete(id uint) error
}

type candidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) CandidateRepository {
	return &candidateRepository{db: db}
}

func (r *candidateRepository) Create(candidate *schemas.Candidate) error {
	return r.db.Create(candidate).Error
}

func (r *candidateRepository) GetByID(id uint) (*schemas.Candidate, error) {
	var candidate schemas.Candidate
	if err := r.db.Preload("User").First(&candidate, id).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) GetByIDForUser(id, userID uint) (*schemas.Candidate, error) {
	var candidate schemas.Candidate
	if err := r.db.Preload("User").Where("id = ? AND user_id = ?", id, userID).First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) GetByUserID(userID uint) (*schemas.Candidate, error) {
	var candidate schemas.Candidate
	if err := r.db.Preload("User").Where("user_id = ?", userID).First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) GetByEmail(email string) (*schemas.Candidate, error) {
	var candidate schemas.Candidate
	if err := r.db.Preload("User").Where("email = ?", email).First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) GetAll() ([]schemas.Candidate, error) {
	var candidates []schemas.Candidate
	if err := r.db.Preload("User").Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *candidateRepository) Update(candidate *schemas.Candidate) error {
	return r.db.Save(candidate).Error
}

func (r *candidateRepository) Delete(id uint) error {
	return r.db.Delete(&schemas.Candidate{}, id).Error
}
