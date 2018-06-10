package validator

import (
	"log"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type Validate struct {
	*validator.Validate
}

func New() *Validate {
	v := &Validate{validator.New()}
	mustRegisterValidation(v, "phone", phone)
	mustRegisterValidation(v, "digits", digits)
	mustRegisterValidation(v, "digital_id", digital_id)
	return v
}

func mustRegisterValidation(v *Validate, name string, fn validator.Func) {
	if err := v.RegisterValidation(name, fn); err != nil {
		log.Fatal(err)
	}
}

var phoneRegexp = regexp.MustCompile("^[0-9]{10}$")

// TODO: better phone validation (use lib maybe?)
func phone(fl validator.FieldLevel) bool {
	return phoneRegexp.MatchString(fl.Field().String())
}

var digitsRegexp = regexp.MustCompile("^[0-9]*$")

func digits(fl validator.FieldLevel) bool {
	return digitsRegexp.MatchString(fl.Field().String())
}

var digitsIDRegexp = regexp.MustCompile("^[1-9]+[0-9]*$")

func digital_id(fl validator.FieldLevel) bool {
	return digitsIDRegexp.MatchString(fl.Field().String())
}
