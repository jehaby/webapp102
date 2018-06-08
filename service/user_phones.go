package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jehaby/webapp102/entity"
	uuid "github.com/satori/go.uuid"
)

type CreatePhoneArgs struct {
	CountryCode int64     `json:"country_code" validator:"required,min=1,max=999"`
	Number      string    `json:"number" validator:"required,"`
	UserUUID    uuid.UUID `json:"user_uuid"`
}

func (a CreatePhoneArgs) Phone() entity.Phone {
	return entity.Phone{
		UUID:        uuid.NewV4(),
		CountryCode: a.CountryCode,
		Number:      a.Number,
		UserUUID:    a.UserUUID,
	}
}

func (us *UserService) CreatePhone(args CreatePhoneArgs) (entity.Phone, error) {
	err := us.val.Struct(args)
	if err != nil {
		return entity.Phone{}, err
	}

	phone := args.Phone()

	_, err = us.db.Model(&phone).Returning("*").Insert()
	if err != nil {
		if err, ok := err.(IntegrityViolationer); ok && err.IntegrityViolation() {
			spew.Dump("fucking integrity violation")
			return phone, &ErrBadRequest{err}
		}

		// TODO: err code 400 (unique, integrity violations)
		return phone, err
	}

	return phone, nil
}

func (us *UserService) DeletePhone(uuid uuid.UUID) error {

	return nil
}
