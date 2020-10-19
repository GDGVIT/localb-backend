package admin

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(username string) (*models.Admin, error)

	GetBusinessesToApprove() (*[]models.Business, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}
	err := r.DB.Where("username = ?", username).First(admin).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	if admin.Username == "" {
		return nil, pkg.ErrNotFound
	}
	return admin, nil
}

func (r *repo) GetBusinessesToApprove() (*[]models.Business, error) {
	var bizs []models.Business

	err := r.DB.Where("approved=?", false).Find(&bizs).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}

	return &bizs, nil
}
