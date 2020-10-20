package business

import "github.com/rithikjain/local-businesses-backend/pkg/models"

type Service interface {
	AddBusiness(business *models.Business) error

	GetApprovedBusinesses(page, pageSize int) (*[]models.Business, error)

	GetBusinessesByCity(city string) (*[]models.Business, error)

	GetRepo() Repository
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) AddBusiness(business *models.Business) error {
	return s.repo.AddBusiness(business)
}

func (s *service) GetApprovedBusinesses(page, pageSize int) (*[]models.Business, error) {
	return s.repo.GetApprovedBusinesses(page, pageSize)
}

func (s *service) GetBusinessesByCity(city string) (*[]models.Business, error) {
	return s.repo.GetBusinessesByCity(city)
}

func (s *service) GetRepo() Repository {
	return s.repo
}
