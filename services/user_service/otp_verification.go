package userService

import (
	appErr "amar_dokan/app_error"
	"amar_dokan/models"
	"amar_dokan/utils"
	"time"
)

// OtpVerification implements [UserService].
func (s *userService) OtpVerification(req *models.OtpVerifyRequest) (string, error) {
	if len(req.Otp) != 6 {
		return "", appErr.ErrOTPInvalid
	}

	tuser, err := s.repo.PandingUserFindById(req.Uid)

	if err != nil {
		return "", appErr.ErrInvalidID
	}

	if req.Otp != tuser.Otp {
		return "", appErr.ErrOTPInvalid
	}

	u := &models.User{
		Name:     tuser.Name,
		Email:    tuser.Email,
		Password: tuser.Password,
		IsOwner:  tuser.IsOwner,
	}

	isValid := time.Since(tuser.CreatedAt).Seconds() <= 120
	if !isValid {
		return "", appErr.ErrOTPExpired
	}

	user, err := s.repo.Create(u)
	if err != nil {
		return "", err
	}

	err = s.repo.DeletePandingUser(req.Uid)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user, utils.AccessToken, s.jwtSecret, s.jwtExpiryDays)
	return token, err

}
