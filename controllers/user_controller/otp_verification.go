package usercontroller

import (
	"amar_dokan/models"
	"amar_dokan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OtpVerification godoc
// @Summary Verify OTP
// @Description Verify a newly registered user's OTP and return an access token
// @Tags Users
// @Accept json
// @Produce json
// @Param payload body models.OtpVerifyRequest true "OTP verification payload"
// @Success 200 {object} utils.ApiResponse "Otp Verification Successful"
// @Failure 400 {object} utils.ApiResponse "Invalid request payload"
// @Failure 500 {object} utils.ApiResponse "Internal server error"
// @Router /users/verification [post]
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
