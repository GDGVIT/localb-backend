package admin

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(username, password string) (*models.Admin, error)

	GetBusinessesToApprove(page, pageSize int) (*[]models.Business, error)

	ApproveBusiness(businessID string) error

	DeleteBusiness(businessID string) error

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

func (s *service) Login(username, password string) (*models.Admin, error) {
	admin, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, admin.Password) {
		return admin, nil
	}
	return nil, pkg.ErrNotFound
}

func (s *service) GetBusinessesToApprove(page, pageSize int) (*[]models.Business, error) {
	return s.repo.GetBusinessesToApprove(page, pageSize)
}

func (s *service) ApproveBusiness(businessID string) error {
	return s.repo.ApproveBusiness(businessID)
}

func (s *service) DeleteBusiness(businessID string) error {
	return s.repo.DeleteBusiness(businessID)
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
