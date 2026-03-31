package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RefrashToken godoc
// @Summary Refresh access token
// @Description Generate a new access token using the Authorization bearer token
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer refresh token"
// @Success 200 {object} utils.ApiResponse "Access Token Genaret Successful"
// @Failure 401 {object} utils.ApiResponse "Unauthorized"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users/refresh-token [get]
func (ctrl *userController) RefrashToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Authorization header required",
		})
		c.Abort()
		return
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		c.JSON(http.StatusUnauthorized, utils.ApiResponse{
			Success: false,
			Message: "Format: Bearer <token>",
		})
		c.Abort()
		return
	}

	token, err := ctrl.service.RefreshToken(parts[1])
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Access Token Genaret Successful",
		Data:    models.TokenResponse{Token: token},
	})
}
