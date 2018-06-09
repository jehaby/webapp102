package validator

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

var phoneRegexp = regexp.MustCompile("^[0-9]{10}$")

// TODO: better phone validation (use lib maybe?)
func phone(fl validator.FieldLevel) bool {
	return phoneRegexp.MatchString(fl.Field().String())
}
