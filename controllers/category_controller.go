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
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{service: service}
}

// Create implements [CategoryController].
func (ctr *categoryController) Create(c *gin.Context) {
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

	data, erro := ctr.service.Create(&req)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: "Internal server erro",
		})
	}

	c.JSON(
		http.StatusOK,
		utils.ApiResponse{
			Success: true,
			Message: "Create Successful",
			Data:    data,
		},
	)

}
