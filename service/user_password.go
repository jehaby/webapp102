package service

import (
	"context"
	"time"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/random"
)

const passwordResetTokenLenght = 32

func (us *UserService) ProcessPasswordResetRequest(ctx context.Context, nameOrEmail string) error {
	tkn, err := random.GenerateRandomString(passwordResetTokenLenght)
	if err != nil {
		return err
	}

	_, err = us.db.
		Model(&entity.User{}).
		Set("confirmation_token = ?", tkn).
		Set("confirmation_token_created_at = NOW()").
		Set("updated_at = NOW()").
		Where("name = ?", nameOrEmail).WhereOr("email = ?", nameOrEmail).
		Update()

	return err
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
