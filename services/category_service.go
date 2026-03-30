package services

import (
	appErr "amar_dokan/app_error"
	"amar_dokan/models"
	"amar_dokan/repositories"
	"amar_dokan/utils"
)

type CategoryService interface {
	Create(req *models.CategoryRequest, token string) (*models.Category, error)
	Get() ([]*models.Category, error)
	Delete(cID string, token string) (string, error)
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

func (c categoryService) Create(req *models.CategoryRequest, token string) (*models.Category, error) {

	payload, err := utils.DecodeJWT(token, c.SecureKey)
	if err != nil {
		return nil, appErr.ErrInternalServer
	}

	cdata := models.Category{
		UID:      payload.Sub,
		ImageUrl: req.ImageUrl,
		Name:     req.Name,
	}

	res, error := c.repo.Create(&cdata)

	if error != nil {
		return nil, error
	}

	return res, nil
}

func (c categoryService) Get() ([]*models.Category, error) {
	return c.repo.List()
}

func (c categoryService) Delete(cID string, token string) (string, error) {

	payload, err := utils.DecodeJWT(token, c.SecureKey)
	if err != nil {
		return "", appErr.ErrInternalServer
	}

	return c.repo.Delete(cID, payload.Sub)
}
