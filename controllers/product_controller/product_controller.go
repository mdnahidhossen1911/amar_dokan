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

// Create godoc
// @Summary Create product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body models.ProductRequest true "Product payload"
// @Success 201 {object} utils.ApiResponse "Note Created"
// @Failure 400 {object} utils.ApiResponse "Invalid request payload"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /products [post]
func (p productcontroller) Create(c *gin.Context) {
	var req models.ProductRequest
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

// Delete godoc
// @Summary Delete product
// @Description Delete a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /products/{id} [delete]
func (p productcontroller) Delete(c *gin.Context) {
	id := c.Param("id")

	message, err := p.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: message,
	})
}

// Get godoc
// @Summary List products
// @Description Get all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /products [get]
func (p productcontroller) Get(c *gin.Context) {
	notes, err := p.service.Get()
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

// Update godoc
// @Summary Update product
// @Description Update a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param payload body models.ProductRequest true "Product payload"
// @Success 200 {object} utils.ApiResponse "Product update successfully"
// @Failure 400 {object} utils.ApiResponse "Invalid request payload"
// @Failure 404 {object} utils.ApiResponse "Product not found"
// @Router /products/{id} [put]
func (p productcontroller) Update(c *gin.Context) {

	id := c.Param("id")

	var product models.ProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	pdrt := models.ProductUpdateRequest{
		ID:          id,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
	}

	pdata, err := p.service.Update(pdrt)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Product update successfully",
		Data:    pdata,
	})

}
