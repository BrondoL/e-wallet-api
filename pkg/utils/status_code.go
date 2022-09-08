package utils

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
)

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*custom_error.NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*custom_error.UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	}
	return statusCode
}
