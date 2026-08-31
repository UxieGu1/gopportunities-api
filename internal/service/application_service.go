package service

import (
	"errors"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type ApplicationService interface {
	Create(application *schemas.Application) error
	GetByID(id uint) (*schemas.Application, error)
	GetAll() ([]schemas.Application, error)
	Update(application *schemas.Application) error
	Delete(id uint) error
}

type applicationService struct {
	repo repository.ApplicationRepository
}

func NewApplicationService(repo repository.ApplicationRepository) ApplicationService {
	return &applicationService{repo: repo}
}

func (s *applicationService) Create(application *schemas.Application) error {
	if application.CandidateID == 0 || application.OpeningID == 0 {
		return errors.New("candidate ID e opening ID são obrigatórios")
	}
	if application.Status == "" {
		application.Status = "APPLIED" 
	}
	return s.repo.Create(application)
}

func (s *applicationService) GetByID(id uint) (*schemas.Application, error) {
	return s.repo.GetByID(id)
}

func (s *applicationService) GetAll() ([]schemas.Application, error) {
	return s.repo.GetAll()
}

func (s *applicationService) Update(application *schemas.Application) error {
	return s.repo.Update(application)
}

func (s *applicationService) Delete(id uint) error {
	return s.repo.Delete(id)
}