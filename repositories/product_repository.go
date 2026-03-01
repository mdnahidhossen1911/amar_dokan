package repositories

import (
	"amar_dokan/models"
	"fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(note *models.Product) (*models.Product, error)
	List() ([]*models.Product, error)
	Update(req *models.ProductUpdateRequest) (*models.Product, error)
	Delete(Id string) (string, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{db: db}
}

// Create implements [NoteRepository].
func (p productRepository) Create(note *models.Product) (*models.Product, error) {

	if err := p.db.Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil

}

func (p productRepository) List() ([]*models.Product, error) {
	var notes []*models.Product

	if err := p.db.Where("is_delete = false").Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// Delete implements [NoteRepository].
func (p productRepository) Delete(Id string) (string, error) {
	var product models.Product
	if err := p.db.Where("id = ?", Id).First(&product).Error; err != nil {
		return "", fmt.Errorf("product not found")
	}
	if err := p.db.Model(product).Where("id = ?", Id).Updates(map[string]interface{}{
		"is_delete": true,
	}).Error; err != nil {
		return "", err
	}
	return "Product deleted successfully", nil
}

// Update implements [NoteRepository].
func (p productRepository) Update(req *models.ProductUpdateRequest) (*models.Product, error) {
	var product models.Product

	if err := p.db.Model(product).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"image_url":   req.ImageUrl,
		"price":       req.Price,
	}).Error; err != nil {
		return nil, err
	}

	if err := p.db.Where("id = ?", req.ID).First(&product).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}
	return &product, nil
}
