package utils

import (
	"errors"
	"io"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errorMessage []string

	if !errors.Is(err, io.EOF) {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, e.Field()+" "+e.Tag())
		}
		return errorMessage
	}
	return []string{"No Data Found"}
}
