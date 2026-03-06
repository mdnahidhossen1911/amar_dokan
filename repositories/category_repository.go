package repositories

import (
	appErr "amar_dokan/app_error"
	"amar_dokan/models"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Create(c *models.Category) (*models.Category, error)
	List() ([]*models.Category, error)
	Update(c *models.CategoryRequest) (*models.Category, error)
	Delete(cid, uid string) (string, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}

// Create implements [CategoryRepo].
func (c *categoryRepo) Create(crt *models.Category) (*models.Category, error) {
	if err := c.db.Create(&crt).Error; err != nil {
		return nil, appErr.ErrInternalServer
	}
	return crt, nil
}

func (c *categoryRepo) List() ([]*models.Category, error) {
	var categorys []*models.Category

	if err := c.db.Where("is_delete = false").Find(&categorys).Error; err != nil {
		return nil, appErr.ErrInternalServer
	}
	return categorys, nil
}

// Update implements [CategoryRepo].
func (c *categoryRepo) Update(ctr *models.CategoryRequest) (*models.Category, error) {
	panic("unimplemented")
}

// Delete implements [CategoryRepo].
func (c *categoryRepo) Delete(cid, uid string) (string, error) {

	var ctr *models.Category
	result := c.db.Model(&ctr).Where("uid = ? And id = ?", uid, cid).Updates(map[string]interface{}{
		"is_delete": true,
	})

	if result.Error != nil {
		return "", appErr.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", appErr.ErrCategoryNotFound
	}

	return "Delete category Succesfully", nil

}
