package repositories

import (
	"amar_dokan/models"
	"fmt"

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

// Get implements [AddToCartRepository].
func (a *addToCartRepository) Get(uid string) ([]*models.AddToCart, error) {
	var addToCarts []*models.AddToCart

	err := a.db.Preload("Product").Where("user_id = ? And is_delete = false", uid).Find(&addToCarts).Error
	if err != nil {
		return nil, err
	}
	return addToCarts, nil
}

// Update implements [AddToCartRepository].
func (a *addToCartRepository) Update(addToCart *models.AddToCartUpdateRequest) (*models.AddToCart, error) {

	var existing models.AddToCart

	error := a.db.Model(existing).Where("id = ? And user_id = ?", addToCart.ID, addToCart.UserID).First(&existing).Error
	if error != nil {
		return nil, fmt.Errorf("not found")
	}

	existing.Quantity = addToCart.Quantity
	if err := a.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil

}

// Delete implements [AddToCartRepository].
func (a *addToCartRepository) Delete(addToCart *models.AddToCartUpdateRequest) (string, error) {

	var cart models.AddToCart

	result := a.db.
		Model(&cart).
		Where("id = ? AND user_id = ?", addToCart.ID, addToCart.UserID).
		Update("is_delete", true)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("no cart found with id %s", addToCart.ID)
	}

	return fmt.Sprintf("Cart %s deleted successfully", addToCart.ID), nil
}
