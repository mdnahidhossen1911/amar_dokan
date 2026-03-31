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

// Create godoc
// @Summary Add product to cart
// @Description Add a product to the authenticated user's cart
// @Tags Add To Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body models.AddToCartRequest true "Add to cart payload"
// @Success 200 {object} utils.ApiResponse "Add to cart successfully"
// @Failure 400 {object} utils.ApiResponse "Payload is not valid"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /add-to-cart [post]
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

// Get godoc
// @Summary Get cart items
// @Description Get all cart items for the authenticated user
// @Tags Add To Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /add-to-cart [get]
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

// Update godoc
// @Summary Update cart item quantity
// @Description Update quantity for a specific cart item
// @Tags Add To Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart item ID"
// @Param payload body models.AddToCartQuantityRequest true "Quantity payload"
// @Success 200 {object} utils.ApiResponse "quantity Update Successfully"
// @Failure 400 {object} utils.ApiResponse "Invalid payload"
// @Failure 404 {object} utils.ApiResponse "Cart item not found"
// @Router /add-to-cart/{id} [put]
func (a *addToCardController) Update(c *gin.Context) {

	id := c.Param("id")

	token := utils.GetTokenFromHeader(c)

	var req models.AddToCartQuantityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: "Invalid payload, 'quantity' is requird",
		})
		return
	}

	data := models.AddToCartUpdateRequest{
		ID:       id,
		Quantity: req.Quantity,
	}

	atc, err := a.service.Update(&data, token)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Message: "Not Found this card",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "quantity Update Successfully",
		Data:    atc,
	})

}

// Delete godoc
// @Summary Delete cart item
// @Description Delete a specific item from the authenticated user's cart
// @Tags Add To Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart item ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse "Cart item not found"
// @Router /add-to-cart/{id} [delete]
func (a *addToCardController) Delete(c *gin.Context) {
	id := c.Param("id")

	token := utils.GetTokenFromHeader(c)

	res, erro := a.service.Delete(id, token)

	if erro != nil {
		c.JSON(
			http.StatusNotFound, utils.ApiResponse{
				Success: false,
				Message: erro.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: res,
	})
}
