package utils

import "github.com/go-playground/validator/v10"

func ValidatorErrors(err error) map[string]string {

	fields := map[string]string{}

	if _, ok := err.(validator.ValidationErrors); ok {
		for _, err := range err.(validator.ValidationErrors) {
			fields[err.Field()] = err.Error()
		}
		return fields
	}

	fields = map[string]string{
		"error": err.Error(),
	}

	return fields
}
