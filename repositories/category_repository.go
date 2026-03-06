package repositories

import (
	appErr "amar_dokan/app_error"
	"amar_dokan/models"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(c *models.Category) (*models.Category, error)
	List() ([]*models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

// Create implements [CategoryRepo].
func (c *categoryRepo) Create(crt *models.Category) (*models.Category, error) {
	if err := c.db.Create(&crt).Error; err != nil {
		return nil, appErr.ErrInternalServer
	}
	return crt, nil
}

func (c *categoryRepo) List() ([]*models.Category, error) {
	var categorys []*models.Category

	if err := c.db.Where("is_delete = false").Find(&categorys).Error; err != nil {
		return nil, appErr.ErrInternalServer
	}
	return categorys, nil
}
