package service

import (
	"errors"

	"github.com/UxieGu1/gopportunities-api/internal/repository"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
)

type CompanyService interface {
	Create(company *schemas.Company) (*schemas.CompanyResponse, error)
	GetByID(id uint) (*schemas.CompanyResponse, error)
	GetByIDForUser(id, userID uint) (*schemas.CompanyResponse, error)
	GetAll() ([]schemas.CompanyResponse, error)
	Update(id uint, companyData *schemas.Company) (*schemas.CompanyResponse, error)
	UpdateForUser(id, userID uint, companyData *schemas.Company) (*schemas.CompanyResponse, error)
	Delete(id uint) error
	DeleteForUser(id, userID uint) error
}

type companyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) CompanyService {
	return &companyService{repo: repo}
}

func mapToResponse(c *schemas.Company) *schemas.CompanyResponse {
	return &schemas.CompanyResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Website:     c.Website,
		Email:       c.Email,
		Location:    c.Location,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (s *companyService) Create(company *schemas.Company) (*schemas.CompanyResponse, error) {

	if company.Name == "" {
		return nil, errors.New("o nome da empresa é obrigatório")
	}

	if err := s.repo.Create(company); err != nil {
		return nil, err
	}

	return mapToResponse(company), nil
}

func (s *companyService) GetByID(id uint) (*schemas.CompanyResponse, error) {
	company, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mapToResponse(company), nil
}

func (s *companyService) GetByIDForUser(id, userID uint) (*schemas.CompanyResponse, error) {
	company, err := s.repo.GetByID(id)
	if err != nil || company.UserID != userID {
		return nil, errors.New("empresa não encontrada")
	}
	return mapToResponse(company), nil
}

func (s *companyService) GetAll() ([]schemas.CompanyResponse, error) {
	companies, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]schemas.CompanyResponse, 0, len(companies))
	for _, c := range companies {
		responses = append(responses, *mapToResponse(&c))
	}

	return responses, nil
}

func (s *companyService) Update(id uint, companyData *schemas.Company) (*schemas.CompanyResponse, error) {
	company, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("empresa não encontrada")
	}

	company.Name = companyData.Name
	company.Description = companyData.Description
	company.Website = companyData.Website
	company.Email = companyData.Email
	company.Location = companyData.Location

	if err := s.repo.Update(company); err != nil {
		return nil, err
	}

	return mapToResponse(company), nil
}

func (s *companyService) UpdateForUser(id, userID uint, companyData *schemas.Company) (*schemas.CompanyResponse, error) {
	company, err := s.repo.GetByID(id)
	if err != nil || company.UserID != userID {
		return nil, errors.New("empresa não encontrada")
	}
	company.Name = companyData.Name
	company.Description = companyData.Description
	company.Website = companyData.Website
	company.Email = companyData.Email
	company.Location = companyData.Location
	if err := s.repo.Update(company); err != nil {
		return nil, err
	}
	return mapToResponse(company), nil
}

func (s *companyService) Delete(id uint) error {
	if _, err := s.repo.GetByID(id); err != nil {
		return errors.New("empresa não encontrada")
	}

	return s.repo.Delete(id)
}

func (s *companyService) DeleteForUser(id, userID uint) error {
	company, err := s.repo.GetByID(id)
	if err != nil || company.UserID != userID {
		return errors.New("empresa não encontrada")
	}
	return s.repo.Delete(id)
}
