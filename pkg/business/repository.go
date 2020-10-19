package business

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	AddBusiness(business *models.Business) error
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) AddBusiness(business *models.Business) error {
	err := r.DB.Create(business).Error
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}
