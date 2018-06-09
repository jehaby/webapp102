package validator

import (
	"log"

	"gopkg.in/go-playground/validator.v9"
)

type Validate struct {
	*validator.Validate
}

func New() *Validate {
	res := &Validate{validator.New()}
	err := res.RegisterValidation("phone", phone)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
