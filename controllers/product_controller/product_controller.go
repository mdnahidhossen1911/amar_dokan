package productcontroller

import (
	"amar_dokan/models"
	productservice "amar_dokan/services/product_service"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func NewProductController(srv productservice.ProductService) ProductController {
	return productcontroller{
		service: srv,
	}
}

type productcontroller struct {
	service productservice.ProductService
}

func (p productcontroller) Create(c *gin.Context) {
	var req models.NoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token := utils.GetTokenFromHeader(c)

	note, error := p.service.Create(&req, token)
	if error != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: error.Error(),
		})
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Note Created",
		Data:    note,
	})

}

// Delete implements [NoteController].
func (p productcontroller) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Get implements [NoteController].
func (p productcontroller) Get(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)

	notes, err := p.service.Get(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    notes,
	})

}

// GetById implements [NoteController].
func (p productcontroller) GetById(c *gin.Context) {
	panic("unimplemented")
}

// Update implements [NoteController].
func (p productcontroller) Update(c *gin.Context) {
	panic("unimplemented")
}
