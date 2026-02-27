package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *userController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	token, err := ctrl.service.Login(&req)
	if err != nil {
		if err == models.ErrUserNotFound || err == models.ErrInvalidPassword {
			c.JSON(http.StatusUnauthorized, utils.ApiResponse{
				Success: false,
				Message: "Invalid credentials",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Login Successful",
		Data:    *token,
	})
}
