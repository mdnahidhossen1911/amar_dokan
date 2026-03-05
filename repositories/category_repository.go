package repositories

import (
	"amar_dokan/models"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(c *models.Category) (*models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}
// Create implements [CategoryRepo].
func (*categoryRepo) Create(c *models.Category) (*models.Category, error) {
	panic("unimplemented")
}


