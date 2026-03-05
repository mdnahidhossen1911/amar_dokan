package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctrl *userController) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		if strings.Contains(err.Error(), "failed on the 'email' tag") {
			c.JSON(http.StatusBadRequest, utils.ApiResponse{
				Success: false,
				Message: "Invalid email address.",
			})
			return
		}

		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := ctrl.service.Register(&req)
	if err != nil {
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusCreated, utils.ApiResponse{
		Success: true,
		Message: "Account created. OTP has been sent to your email.",
		Data:    *user,
	})

}
