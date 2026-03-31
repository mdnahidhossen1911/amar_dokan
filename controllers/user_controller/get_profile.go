package usercontroller

import (
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Get the authenticated user's profile
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users/profile [get]
func (ctrl *userController) GetProfile(c *gin.Context) {

	token := utils.GetTokenFromHeader(c)

	user, err := ctrl.service.GetProfile(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Data:    "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    user,
	})

}
