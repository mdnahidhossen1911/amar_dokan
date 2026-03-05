package controllers

import (
	"amar_dokan/services"

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
	panic("unimplemented")
}
