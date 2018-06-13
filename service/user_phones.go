package service

import (
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
)

type CreatePhoneArgs struct {
	CountryCode int64     `json:"country_code" validate:"required,min=1,max=999"`
	Number      string    `json:"number" validate:"required,phone"`
	UserUUID    uuid.UUID `json:"user_uuid" validate:"required"`
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
		return entity.Phone{}, &ErrBadRequest{err}
	}

	phone := args.Phone()

	err = us.db.RunInTransaction(func(tx *pg.Tx) error {
		_, err = tx.Model(&phone).Returning("*").Insert()
		if err != nil {
			if err, ok := err.(IntegrityViolationer); ok && err.IntegrityViolation() {
				return &ErrBadRequest{err}
			}
			// TODO: err code 400 (unique, integrity violations)
			return err
		}

		// setting newly created phone to default phone if not alredy set
		_, err = tx.Model(&entity.User{UUID: args.UserUUID}).
			Set("default_phone = ?", phone.UUID).
			Where("uuid = ?", args.UserUUID).
			Where("default_phone IS NULL").
			Update()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		return nil
	})
	if err != nil {
		return phone, err
	}

	return phone, nil
}

type DeletePhoneArgs struct {
	UserUUID, PhoneUUID uuid.UUID
}

func (us *UserService) DeletePhone(args DeletePhoneArgs) (entity.User, error) {
	user := entity.User{}

	err := us.db.RunInTransaction(func(tx *pg.Tx) error {

		// check if phone is used in active ads; return error if it is.
		// TODO: optimise query? (use WITH or JOIN)?
		activeAdsCnt, err := tx.Model(&entity.Ad{}).
			Where("user_uuid = ?", args.UserUUID).
			Where("phone_uuid = ?", args.PhoneUUID).
			Where("deleted_at IS NULL").
			Count()
		if err != nil {
			return err
		} else if activeAdsCnt != 0 {
			// TODO: error!
			return &ErrBadRequest{errors.New("err_phone_used_in_active_ads")}
		}

		// setting user's default phone for any other user's phone
		_, err = tx.Model(&user).
			Set("default_phone = (SELECT uuid FROM phones WHERE user_uuid = ? AND uuid != ? LIMIT 1)", args.UserUUID, args.PhoneUUID).
			Where("uuid = ?", args.UserUUID).
			Returning("*").
			Update()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		err = tx.Delete(&entity.Phone{UUID: args.PhoneUUID})
		if err != nil {
			// TODO: better err check
			if err == pg.ErrNoRows {
				return ErrNotFound
			}
			return err
		}

		return nil
	})

	return user, err
}
