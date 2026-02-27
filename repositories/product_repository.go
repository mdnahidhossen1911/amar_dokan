package repositories

import (
	"amar_dokan/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	Create(note *models.Note) (*models.Note, error)
	List(UId string) ([]*models.Note, error)
	Update(req *models.NoteUpdateRequest) (*models.Note, error)
	Delete(Id string) (string, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{db: db}
}

// Create implements [NoteRepository].
func (p productRepository) Create(note *models.Note) (*models.Note, error) {

	if err := p.db.Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil

}

// List implements [NoteRepository].
func (p productRepository) List(UId string) ([]*models.Note, error) {
	var notes []*models.Note

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
func (p productRepository) Update(req *models.NoteUpdateRequest) (*models.Note, error) {
	panic("unimplemented")
}
