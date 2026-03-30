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

// Create implements [CategoryController].
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
	if err != nil{
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: res,
	})
}
