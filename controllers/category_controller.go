package controllers

import (
	"amar_dokan/models"
	"amar_dokan/services"
	"amar_dokan/utils"
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
	return categoryController{service: service}
}

// Create implements [CategoryController].
func (categoryController) Create(c *gin.Context) {
	var req models.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,
			utils.ApiResponse{
				Success: false,
				Message: "Invalid Payload"})
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

	

}
