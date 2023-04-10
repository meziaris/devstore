package internal

import (
	"net/http"

	"github.com/meziaris/devstore/internal/app/schema"
)

func APIResponse(code int, message string, data ...interface{}) schema.ApiResponse {
	response := schema.ApiResponse{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
		Data:    data,
	}
	return response
}
