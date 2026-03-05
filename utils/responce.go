package utils

import (
	appErr "amar_dokan/app_error"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func ErrorResponce(err error) (code int, obj any) {

	if appErr, ok := err.(*appErr.Error); ok {
		return appErr.Status, ApiResponse{Success: false, Message: appErr.Message}
	}

	return 500, ApiResponse{Success: false, Message: "Internal Server Error"}
}
