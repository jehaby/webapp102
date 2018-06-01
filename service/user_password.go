package service

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/random"
)

func (us *UserService) ProcessPasswordResetRequest(ctx context.Context, nameOrEmail string) error {
	user, err := us.GetByNameOrEmail(nameOrEmail)
	if err != nil {
		return ignorePgNotFoundErr(err)
	}

	tkn, err := random.GenerateRandomString(32)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, UserCtxKey, &user)
	_, err = us.Update(ctx, user.UUID, UserUpdateArgs{
		ConfirmationToken: pointer.ToString(tkn),
	})
	if err != nil {
		return err
	}

	// TODO: send email
	return nil
}

func (us *UserService) ProcessPasswordResetAction(ctx context.Context, token, newPassword string) (entity.User, error) {
	user := entity.User{}
	err := us.db.Model(&user).Where("confirmation_token = ?", token).First()
	if err != nil {
		return user, checkPgNotFoundErr(err)
	}

	// TODO: check for bugs with timezones
	if time.Since(user.ConfirmationTokenCreatedAt) > time.Hour {
		return user, ErrTokenExpired
	}

	user, err = us.ChangePassword(ctx, user.UUID, newPassword)
	if err != nil {
		return user, err
	}

	return user, nil
}
