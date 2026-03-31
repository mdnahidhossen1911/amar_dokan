package usercontroller

import (
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetByID godoc
// @Summary Get user by ID
// @Description Get a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Router /users/{id} [get]
func (ctrl *userController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Data:    user,
	})
}
