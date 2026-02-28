package repositories

import (
	"amar_dokan/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(note *models.Product) (*models.Product, error)
	List(UId string) ([]*models.Product, error)
	Update(req *models.NoteUpdateRequest) (*models.Product, error)
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

// List implements [NoteRepository].
func (p productRepository) List(UId string) ([]*models.Product, error) {
	var notes []*models.Product

	if err := p.db.Where("uid = ?", UId).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// Delete implements [NoteRepository].
func (p productRepository) Delete(Id string) (string, error) {
	panic("unimplemented")
}

// Update implements [NoteRepository].
func (p productRepository) Update(req *models.NoteUpdateRequest) (*models.Product, error) {
	panic("unimplemented")
}
