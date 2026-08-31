package service

import (
	"errors"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type CandidateService interface {
	Create(candidate *schemas.Candidate) error
	GetByID(id uint) (*schemas.Candidate, error)
	GetByEmail(email string) (*schemas.Candidate, error)
	GetAll() ([]schemas.Candidate, error)
	Update(candidate *schemas.Candidate) error
	Delete(id uint) error
}

type candidateService struct {
	repo repository.CandidateRepository
}

func NewCandidateService(repo repository.CandidateRepository) CandidateService {
	return &candidateService{repo: repo}
}

func (s *candidateService) Create(candidate *schemas.Candidate) error {
	if candidate.User.Email == "" {
		return errors.New("o email do usuário associado é obrigatório")
	}
	return s.repo.Create(candidate)
}

func (s *candidateService) GetByID(id uint) (*schemas.Candidate, error) {
	return s.repo.GetByID(id)
}

func (s *candidateService) GetByEmail(email string) (*schemas.Candidate, error) {
	return s.repo.GetByEmail(email)
}

func (s *candidateService) GetAll() ([]schemas.Candidate, error) {
	return s.repo.GetAll()
}

func (s *candidateService) Update(candidate *schemas.Candidate) error {
	return s.repo.Update(candidate)
}

func (s *candidateService) Delete(id uint) error {
	return s.repo.Delete(id)
}