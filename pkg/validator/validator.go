package validator

import (
	"fmt"
	"microserviceMOCK/internal/core/domain"
	"net/http"
	"time"

	validators "github.com/go-playground/validator/v10"
)

type ErrorOther []error

func (c *ErrorOther) Add(e error) { *c = append(*c, e) }

func (c *ErrorOther) Error() domain.Status {
	var messages []string
	for _, e := range *c {
		messages = append(messages, e.Error())
	}

	return domain.Status{
		Code:    http.StatusBadRequest,
		Message: messages,
	}
}

type Validator interface {
	ValidateStruct(inf interface{}) error
	GenValidateStructErrorMessage(err error) domain.Status
}

type validator struct {
	validator *validators.Validate
}

var (
	dateFormat = "2006-01-02 15:04:05"
)

func New() Validator {
	v := validators.New()
	dateValidator(v)
	return &validator{
		validator: v,
	}
}

func (v *validator) ValidateStruct(inf interface{}) error {

	return v.validator.Struct(inf)
}

func (v *validator) GenValidateStructErrorMessage(err error) domain.Status {
	var messages []string
	validationErrors := err.(validators.ValidationErrors)
	for _, e := range validationErrors {
		if e.Tag() == "max" {
			messages = append(messages, fmt.Sprintf("%s::Sorry, %s are limited to %s characters", e.Field(), e.Field(), e.Param()))
		} else if e.Tag() == "numeric" {
			messages = append(messages, fmt.Sprintf("%s::Sorry, %s is wrong number", e.Field(), e.Field()))
		} else if e.Tag() == "oneof" {
			messages = append(messages, fmt.Sprintf("%s::Sorry, %s is not contains '%s'", e.Field(), e.Field(), e.Param()))
		} else if e.Tag() == "multipleLangJSON" {
			messages = append(messages, fmt.Sprintf("%s::Sorry, %s is wrong JSON MultiLanguage", e.Field(), e.Field()))
		} else if e.Tag() == "lte" {
			messages = append(messages, fmt.Sprintf("%s::Sorry, the %s must be less than or equal to %s character", e.Field(), e.Field(), e.Param()))
		} else {
			messages = append(messages, e.Error())
		}
	}

	return domain.Status{
		Code:    http.StatusBadRequest,
		Message: messages,
	}
}

// date format validator ...
func dateValidator(v *validators.Validate) {
	v.RegisterValidation("dateFormat", func(fl validators.FieldLevel) bool {
		if fl != nil {
			d := fl.Field().String()
			_, err := time.Parse(dateFormat, d)
			return err == nil
		} else {
			return true
		}
	})
}
