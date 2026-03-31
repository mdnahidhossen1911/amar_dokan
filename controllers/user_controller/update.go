package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update godoc
// @Summary Update user
// @Description Update a user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param payload body models.User true "User payload"
// @Success 200 {object} models.User
// @Failure 400 {object} utils.ApiResponse "Invalid request payload"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users/{id} [put]
func (ctrl *userController) Update(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	updated, err := ctrl.service.Update(id, &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updated)
}
