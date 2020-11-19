package admin

import (
	"github.com/rithikjain/local-businesses-backend/pkg"
	"github.com/rithikjain/local-businesses-backend/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByUsername(username string) (*models.Admin, error)

	GetBusinessesToApprove(page, pageSize int) (*[]models.Business, error)

	ApproveBusiness(businessID string) error

	DeleteBusiness(businessID string) error
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

func (r *repo) GetBusinessesToApprove(page, pageSize int) (*[]models.Business, error) {
	var bizs []models.Business

	err := r.DB.Where("approved=?", false).Scopes(pkg.Paginate(page, pageSize)).Find(&bizs).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}

	return &bizs, nil
}

func (r *repo) ApproveBusiness(businessID string) error {
	biz := &models.Business{}

	err := r.DB.Where("id = ?", businessID).First(biz).Error
	if err != nil {
		return pkg.ErrDatabase
	}

	err = r.DB.Model(biz).Update("approved", true).Error
	if err != nil {
		return pkg.ErrDatabase
	}

	return nil
}

func (r *repo) DeleteBusiness(businessID string) error {
	biz := &models.Business{}

	err := r.DB.Where("id = ?", businessID).First(biz).Error
	if err != nil {
		return pkg.ErrDatabase
	}

	err = r.DB.Unscoped().Delete(biz).Error
	if err != nil {
		return pkg.ErrDatabase
	}

	return nil
}
