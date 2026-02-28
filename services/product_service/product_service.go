package noteservice

import (
	"amar_dokan/models"
	"amar_dokan/repositories"
	"amar_dokan/utils"
)

type ProductService interface {
	Create(product *models.ProductRequest, token string) (*models.Product, error)
	Get(token string) ([]*models.Product, error)
	GetProfile(token string) (*models.Product, error)
	Update(models.Product) (*models.Product, error)
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

// Create implements [NoteService].
func (p productService) Create(product *models.ProductRequest, token string) (*models.Product, error) {

	payload, err := utils.DecodeJWT(token, p.jwtSecret)

	if err != nil {
		return nil, err
	}

	noteData := &models.Product{
		UID:   payload.Sub,
		Name: product.Name,
		Description:  product.Description,
		ImageUrl: product.ImageUrl,
		Price: product.Price,
	}

	return p.repo.Create(noteData)

}

func (p productService) Get(token string) ([]*models.Product, error) {

	payload, err := utils.DecodeJWT(token, p.jwtSecret)
	if err != nil {
		return nil, err
	}

	return p.repo.List(payload.Sub)
}

// Delete implements [NoteService].
func (p productService) Delete(id string) (string, error) {
	panic("unimplemented")
}

// GetProfile implements [NoteService].
func (p productService) GetProfile(token string) (*models.Product, error) {
	panic("unimplemented")
}

// Update implements [NoteService].
func (p productService) Update(models.Product) (*models.Product, error) {
	panic("unimplemented")
}
