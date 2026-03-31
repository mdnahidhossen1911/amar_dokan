package controllers

import (
	"amar_dokan/models"
	"amar_dokan/services"
	"amar_dokan/utils"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{service: service}
}

// Create godoc
// @Summary Create category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param payload body models.CategoryRequest true "Category payload"
// @Success 201 {object} utils.ApiResponse "Category create Successful"
// @Failure 400 {object} utils.ApiResponse "Invalid request"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /category [post]
func (ctr *categoryController) Create(c *gin.Context) {

	token := utils.GetTokenFromHeader(c)

	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(utils.ErrorResponce(errors.ErrUnsupported))
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest,
			utils.ApiResponse{
				Success: false,
				Message: "Name is required"})
		return
	}

	if req.ImageUrl == "" {
		c.JSON(http.StatusBadRequest,
			utils.ApiResponse{
				Success: false,
				Message: "Image URL is required"})
		return
	}

	data, erro := ctr.service.Create(&req, token)

	if erro != nil {
		c.JSON(utils.ErrorResponce(erro))
		return
	}

	c.JSON(
		http.StatusCreated,
		utils.ApiResponse{
			Success: true,
			Message: "Category create Successful",
			Data:    data,
		},
	)

}

// Get godoc
// @Summary List categories
// @Description Get all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /category [get]
func (ctr *categoryController) Get(c *gin.Context) {

	data, err := ctr.service.Get()

	if err != nil {
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    data,
	})

}

// Delete implements [CategoryController].
func (ctr *categoryController) Delete(c *gin.Context) {

	id := c.Param("id")
	token := utils.GetTokenFromHeader(c)

	res, err := ctr.service.Delete(id, token)
	if err != nil {
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: res,
	})
}
