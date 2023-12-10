package validators

import (
	"github.com/go-playground/validator"
)

func Validate(data interface{}) (bool, map[string]string) {
	errors := make(map[string]string)
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorField := err.Field()
			errors[errorField] = validateErrorMessage(err)
		}
		return false, errors
	}
	return true, nil
}

func validateErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "min":
		return err.Field() + " minimum length is " + err.Param()
	case "max":
		return err.Field() + " maximum length is " + err.Param()
	case "email":
		return err.Field() + " must be a valid email address"
	default:
		return err.Field() + " is not valid"
	}
}
