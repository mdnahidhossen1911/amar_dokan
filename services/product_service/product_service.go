package noteservice

import (
	"amar_dokan/models"
	"amar_dokan/repositories"
	"amar_dokan/utils"
)

type ProductService interface {
	Create(note *models.NoteRequest, token string) (*models.Note, error)
	Get(token string) ([]*models.Note, error)
	GetProfile(token string) (*models.Note, error)
	Update(models.Note) (*models.Note, error)
	Delete(id string) (string, error)
}

type noteService struct {
	repo      repositories.ProductRepository
	jwtSecret string
}

func NewProductService(key string, repo repositories.ProductRepository) ProductService {
	return noteService{
		jwtSecret: key,
		repo:      repo,
	}
}

// Create implements [NoteService].
func (p noteService) Create(note *models.NoteRequest, token string) (*models.Note, error) {

	payload, err := utils.DecodeJWT(token, p.jwtSecret)

	if err != nil {
		return nil, err
	}

	noteData := &models.Note{
		UID:   payload.Sub,
		Title: note.Title,
		Body:  note.Body,
	}

	return p.repo.Create(noteData)

}

// Get implements [NoteService].
func (p noteService) Get(token string) ([]*models.Note, error) {

	payload, err := utils.DecodeJWT(token, p.jwtSecret)
	if err != nil {
		return nil, err
	}

	return p.repo.List(payload.Sub)
}

// Delete implements [NoteService].
func (p noteService) Delete(id string) (string, error) {
	panic("unimplemented")
}

// GetProfile implements [NoteService].
func (p noteService) GetProfile(token string) (*models.Note, error) {
	panic("unimplemented")
}

// Update implements [NoteService].
func (p noteService) Update(models.Note) (*models.Note, error) {
	panic("unimplemented")
}
