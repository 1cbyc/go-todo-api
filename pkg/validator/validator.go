package validator

import (
	"github.com/go-playground/validator/v10"
)

// New creates a new validator instance
func New() *validator.Validate {
	validate := validator.New()
	return validate
}

// Global validator instance
var Validate = New()
