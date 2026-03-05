package services

import (
	"amar_dokan/models"
	"amar_dokan/repositories"
)

type CategoryService interface {
	Create(req *models.CategoryRequest) (*models.Category, error)
}

type categoryService struct {
	repo      repositories.CategoryRepo
	SecureKey string
}


func NewCategoryService(repo *repositories.CategoryRepo, secureKey string) CategoryService {
	return categoryService{
		repo:      *repo,
		SecureKey: secureKey,
	}
}

// Create implements [CategoryService].
func (c categoryService) Create(req *models.CategoryRequest) (*models.Category, error) {
	panic("unimplemented")
}
