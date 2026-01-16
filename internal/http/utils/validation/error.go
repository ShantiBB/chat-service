package validation

import (
	"errors"

	ozzovalidation "github.com/go-ozzo/ozzo-validation"
)

//easyjson:json
type ValidateErrorResponse struct {
	Fields map[string]string `json:"fields"`
	Error  string            `json:"error"`
}

func Error(err error) ValidateErrorResponse {
	var validationErrs ozzovalidation.Errors
	errors.As(err, &validationErrs)

	fieldErrors := make(map[string]string)
	for field, fieldErr := range validationErrs {
		fieldErrors[field] = fieldErr.Error()
	}

	response := ValidateErrorResponse{
		Error:  "validation failed",
		Fields: fieldErrors,
	}

	return response
}
