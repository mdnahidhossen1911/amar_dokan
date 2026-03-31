package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login user
// @Description authenticate user and get token
// @Tags Users
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login Credentials"
// @Success 200 {object} utils.ApiResponse "Login Successful"
// @Failure 400 {object} utils.ApiResponse "Invalid request payload"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users/login [post]
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
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Login Successful",
		Data:    *token,
	})
}
