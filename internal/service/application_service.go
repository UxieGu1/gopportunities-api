package service

import (
	"errors"
	"strings"

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
	candidateExists, err := s.repo.CandidateExists(application.CandidateID)
	if err != nil {
		return err
	}
	if !candidateExists {
		return errors.New("candidato não encontrado")
	}
	openingExists, err := s.repo.OpeningExists(application.OpeningID)
	if err != nil {
		return err
	}
	if !openingExists {
		return errors.New("vaga não encontrada")
	}
	openingStatus, err := s.repo.GetOpeningStatus(application.OpeningID)
	if err != nil {
		return err
	}
	if openingStatus != "OPEN" {
		return errors.New("esta vaga não está aceitando candidaturas")
	}
	applicationExists, err := s.repo.Exists(application.CandidateID, application.OpeningID)
	if err != nil {
		return err
	}
	if applicationExists {
		return errors.New("candidato já se aplicou para esta vaga")
	}
	if application.Status == "" {
		application.Status = "APPLIED"
	}
	application.Status = strings.ToUpper(strings.TrimSpace(application.Status))
	if application.Status != "APPLIED" {
		return errors.New("uma nova candidatura deve iniciar com status APPLIED")
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
	application.Status = strings.ToUpper(strings.TrimSpace(application.Status))
	if !validApplicationStatus(application.Status) {
		return errors.New("status de candidatura inválido")
	}

	previous, err := s.repo.GetByID(application.ID)
	if err != nil {
		return err
	}
	if !validStatusTransition(previous.Status, application.Status) {
		return errors.New("transição de status de candidatura inválida")
	}
	return s.repo.Update(application)
}

func validApplicationStatus(status string) bool {
	switch status {
	case "APPLIED", "REVIEWING", "INTERVIEW", "REJECTED", "HIRED":
		return true
	default:
		return false
	}
}

func validStatusTransition(current, next string) bool {
	if current == next {
		return true
	}
	switch current {
	case "APPLIED":
		return next == "REVIEWING" || next == "REJECTED"
	case "REVIEWING":
		return next == "INTERVIEW" || next == "REJECTED"
	case "INTERVIEW":
		return next == "HIRED" || next == "REJECTED"
	default:
		return false
	}
}

func (s *applicationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
