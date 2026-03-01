package noteservice

import (
	"amar_dokan/models"
	"amar_dokan/repositories"
	"amar_dokan/utils"
)

type ProductService interface {
	Create(product *models.ProductRequest, token string) (*models.Product, error)
	Get() ([]*models.Product, error)
	Update(models.ProductUpdateRequest) (*models.Product, error)
	Delete(id string) (string, error)
}

type productService struct {
	repo      repositories.ProductRepository
	jwtSecret string
}

func NewProductService(key string, repo repositories.ProductRepository) ProductService {
	return productService{
		jwtSecret: key,
		repo:      repo,
	}
}

func (p productService) Create(product *models.ProductRequest, token string) (*models.Product, error) {

	payload, err := utils.DecodeJWT(token, p.jwtSecret)

	if err != nil {
		return nil, err
	}

	productData := &models.Product{
		UID:         payload.Sub,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
	}

	return p.repo.Create(productData)

}

func (p productService) Get() ([]*models.Product, error) {

	return p.repo.List()
}

func (p productService) Delete(id string) (string, error) {
	return p.repo.Delete(id)
}

func (p productService) Update(req models.ProductUpdateRequest) (*models.Product, error) {

	return p.repo.Update(&req)

}
