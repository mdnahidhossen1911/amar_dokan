package repositories

import (
	"amar_dokan/models"

	"gorm.io/gorm"
)

type AddToCartRepository interface {
	Create(addToCart *models.AddToCart) (*models.AddToCart, error)
	Get(uid string) ([]*models.AddToCart, error)
	Update(addToCart *models.AddToCartUpdateRequest) (*models.AddToCart, error)
	Delete(addToCart *models.AddToCartUpdateRequest) (string, error)
}

type addToCartRepository struct {
	db *gorm.DB
}

func NewAddToCartRepository(db *gorm.DB) AddToCartRepository {
	return &addToCartRepository{db: db}
}

// Create implements [AddToCartRepository].
func (a *addToCartRepository) Create(addToCart *models.AddToCart) (*models.AddToCart, error) {

	var existing models.AddToCart

	error := a.db.Model(addToCart).Where("product_id = ? And user_id = ?", addToCart.ProductID, addToCart.UserID).First(&existing).Error

	if error == nil {
		existing.Quantity += addToCart.Quantity
		if err := a.db.Save(&existing).Error; err != nil {
			return nil, err
		}

		return &existing, nil
	}

	err := a.db.Create(addToCart).Error
	if err != nil {
		return nil, err
	}
	return addToCart, nil
}

// Delete implements [AddToCartRepository].
func (a *addToCartRepository) Delete(addToCart *models.AddToCartUpdateRequest) (string, error) {
	panic("unimplemented")
}

// Get implements [AddToCartRepository].
func (a *addToCartRepository) Get(uid string) ([]*models.AddToCart, error) {
	var addToCarts []*models.AddToCart

	err := a.db.Preload("Product").Where("user_id = ?", uid).Find(&addToCarts).Error
	if err != nil {
		return nil, err
	}
	return addToCarts, nil
}

// Update implements [AddToCartRepository].
func (a *addToCartRepository) Update(addToCart *models.AddToCartUpdateRequest) (*models.AddToCart, error) {
	panic("unimplemented")
}
