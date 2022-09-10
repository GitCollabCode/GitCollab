package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Validator struct {
	validate *validator.Validate
}

type ValidatorError struct {
	validator.FieldError
}

type ValidatorErrorList struct {
	Messages []string `json:"messages"`
}

type ValidatorErrors []ValidatorError

func (v Validator) Validate(i interface{}) ValidatorErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidatorError
	for _, err := range errs {
		ve := ValidatorError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

func NewValidtor() *Validator {
	validate := validator.New()
	return &Validator{validate}
}

func (v ValidatorError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

func (v ValidatorErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}
