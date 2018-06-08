package validator

import "gopkg.in/go-playground/validator.v9"

type Validate struct {
	*validator.Validate
}

func New() *Validate {
	res := &Validate{validator.New()}

	return res
}
