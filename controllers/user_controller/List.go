package usercontroller

import (
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// List godoc
// @Summary List users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users [get]
func (ctrl *userController) List(c *gin.Context) {
	users, err := ctrl.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    users,
	})
}
