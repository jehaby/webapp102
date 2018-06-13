package service

import (
	"context"
	"fmt"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/random"
)

const (
	passwordResetTokenLenght = 32
	confirmationTokenLife    = "1 hour"
)

var confirmationTokenLifeCond = fmt.Sprintf("confirmation_token_created_at > NOW() - INTERVAL '%s'", confirmationTokenLife)

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

	newPassword, err := hashPassword(newPassword)
	if err != nil {
		return user, err
	}

	res, err := us.db.Model(&user).
		Set("password = ?", newPassword).
		Set("updated_at = NOW()").
		Set("confirmation_token = ?", nil).
		Set("confirmation_token_created_at = ?", nil).
		Where("confirmation_token = ?", token).
		Where(confirmationTokenLifeCond).
		Returning("*").
		Update()

	if err == nil && res != nil && res.RowsAffected() == 0 {
		err = ErrNotFound
	}

	// TODO: think about timezones-related bugs
	return user, err
}
