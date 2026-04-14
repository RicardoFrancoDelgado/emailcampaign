package internallerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationErr := validationErrors[0]

	field := strings.ToLower(validationErr.StructField())
	switch validationErr.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationErr.Param())
	case "min":
		return errors.New(field + " is required " + validationErr.Param())
	case "email":
		return errors.New(field + " is required with max" + validationErr.Param())
	}

	return nil
}
