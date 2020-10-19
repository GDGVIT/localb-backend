package admin

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(username, password string) (*models.Admin, error)

	GetBusinessesToApprove() (*[]models.Business, error)

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

func (s *service) GetBusinessesToApprove() (*[]models.Business, error) {
	return s.repo.GetBusinessesToApprove()
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
