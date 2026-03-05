package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *userController) OtpVerification(c *gin.Context) {
	var req models.OtpVerifyRequest
	if error := c.ShouldBindJSON(&req); error != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Message: error.Error(),
		})
		return
	}

	token, err := ctrl.service.OtpVerification(&req)
	if err != nil {
		c.JSON(utils.ErrorResponce(err))
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{
		Success: true,
		Message: "Otp Verification Successful",
		Data:    models.TokenResponse{Token: token},
	})

}
