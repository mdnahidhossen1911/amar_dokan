package services

import (
	"amar_dokan/models"
	"amar_dokan/repositories"
	"amar_dokan/utils"
)

type AddToCardService interface {
	Create(addToCard *models.AddToCartRequest, token string) (*models.AddToCart, error)
	Get(token string) ([]*models.AddToCart, error)
	Update(addToCard *models.AddToCartUpdateRequest) (*models.AddToCart, error)
	Delete(addToCard *models.AddToCartUpdateRequest) (string, error)
}

type addToCardService struct {
	repo      repositories.AddToCartRepository
	secureKey string
}

func NewAddToCardService(key string, repo repositories.AddToCartRepository) AddToCardService {
	return &addToCardService{secureKey: key, repo: repo}
}

// Create implements [AddToCardService].
func (a *addToCardService) Create(addToCard *models.AddToCartRequest, token string) (*models.AddToCart, error) {

	payload, err := utils.DecodeJWT(token, a.secureKey)
	if err != nil {
		return nil, err
	}

	cartData := &models.AddToCart{
		ProductID: addToCard.ProductID,
		UserID:    payload.Sub,
		Quantity:  1,
	}

	return a.repo.Create(cartData)
}

// Delete implements [AddToCardService].
func (a *addToCardService) Delete(addToCard *models.AddToCartUpdateRequest) (string, error) {
	panic("unimplemented")
}

func (a *addToCardService) Get(token string) ([]*models.AddToCart, error) {
	payload, err := utils.DecodeJWT(token, a.secureKey)
	if err != nil {
		return nil, err
	}
	return a.repo.Get(payload.Sub)
}

// Update implements [AddToCardService].
func (a *addToCardService) Update(addToCard *models.AddToCartUpdateRequest) (*models.AddToCart, error) {
	panic("unimplemented")
}
