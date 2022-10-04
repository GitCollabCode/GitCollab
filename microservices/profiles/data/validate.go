package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError so we do not expose this
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()

	//validate.RegisterValidation("github_username", validateGithubUsername)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// func validateGithubUsername(fl validator.FieldLevel) bool {
// 	re := regexp.MustCompile(`/^[a-z\d](?:[a-z\d]|-(?=[a-z\d])){0,38}$/i`)
// 	sku := re.FindAllString(fl.Field().String(), -1)
// 	return len(sku) == 1
// }
