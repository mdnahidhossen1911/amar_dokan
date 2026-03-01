package controllers

import (
	"amar_dokan/models"
	"amar_dokan/services"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddToCardController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type addToCardController struct {
	service services.AddToCardService
}

func NewAddToCardController(service services.AddToCardService) AddToCardController {
	return &addToCardController{service: service}
}

// Create implements [AddToCardController].
func (a *addToCardController) Create(c *gin.Context) {
	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Payload is not valid",
		})
		return
	}

	token := utils.GetTokenFromHeader(c)

	addToCart, err := a.service.Create(&req, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Add to cart successfully",
		Data:    addToCart,
	})
}

// Delete implements [AddToCardController].
func (a *addToCardController) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Get implements [AddToCardController].
func (a *addToCardController) Get(c *gin.Context) {
	token := utils.GetTokenFromHeader(c)
	addToCarts, err := a.service.Get(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    addToCarts,
	})
}

// Update implements [AddToCardController].
func (a *addToCardController) Update(c *gin.Context) {
	panic("unimplemented")
}
