package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
