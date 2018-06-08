package service

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
	"github.com/jehaby/webapp102/pkg/random"
	"github.com/jehaby/webapp102/pkg/validator"
	"github.com/jehaby/webapp102/service/auth"
)

type UserService struct {
	db  *pg.DB
	val *validator.Validate
	log *log.Logger
}

func newUserService(
	db *pg.DB,
	val *validator.Validate,
	log *log.Logger,
) *UserService {
	return &UserService{
		db:  db,
		val: val,
		log: log,
	}
}

func (us *UserService) GetByNameOrEmail(nameOrEmail string) (entity.User, error) {
	res := entity.User{}
	err := us.db.Model(&res).Where("name = ?", nameOrEmail).WhereOr("email = ?", nameOrEmail).First()
	if err != nil {
		return res, checkPgNotFoundErr(err)
	}
	return res, nil
}

func (us *UserService) GetByUUID(uuid uuid.UUID) (entity.User, error) {
	res := entity.User{}
	err := us.db.Model(&res).
		Relation("Phones").
		Where("uuid = ?", uuid).
		First()
	if err != nil {
		return res, checkPgNotFoundErr(err)
	}
	return res, nil
}

type UserCreateArgs struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3"`
}

const confirmationTokenLenght = 32

func (us *UserService) Create(ctx context.Context, args UserCreateArgs) (entity.User, error) {
	user := entity.User{
		UUID:      uuid.NewV4(),
		CreatedAt: time.Now(),
		Name:      args.Name,
		Email:     args.Email,
	}
	var err error

	user.Password, err = hashPassword(args.Password)
	if err != nil {
		return user, err
	}

	// TODO: send confirmation email
	user.Confirmed = true // TODO: remove in prod
	user.ConfirmationToken, err = random.GenerateRandomString(32)
	if err != nil {
		return user, errors.Wrapf(err, "got error generating token")
	}

	_, err = us.db.Model(&user).Insert()
	if err != nil {
		// TODO: error
		spew.Dump(err)
		return user, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

type UserUpdateArgs struct {
	Email                      *string `validate:"omitempty,email"`
	DefaultPhone               *uuid.UUID
	LastLogout                 *time.Time
	Confirmed                  *bool
	ConfirmationToken          *string
	ConfirmationTokenCreatedAt *time.Time
	BannedAt                   *time.Time
	BannedInfo                 *string
}

func (us *UserService) Update(ctx context.Context, uuid uuid.UUID, args UserUpdateArgs) (*entity.User, error) {
	var (
		user *entity.User
		err  error
	)

	if err = us.val.Struct(args); err != nil {
		return nil, err
	}

	authorizedUser := UserFromCtx(ctx)
	if authorizedUser == nil {
		return nil, ErrNotAuthorized
	}
	if authorizedUser.UUID != uuid {
		if !authorizedUser.IsAdmin() {
			return nil, ErrNotAuthorized
		}
		err = us.db.Model(user).WherePK().First()
		if err != nil {
			return nil, checkPgNotFoundErr(err)
		}
	} else {
		user = authorizedUser
	}

	if args.Email != nil {
		user.Email = *args.Email
	}
	if args.LastLogout != nil {
		user.LastLogout = *args.LastLogout
	}
	if args.Confirmed != nil {
		user.Confirmed = *args.Confirmed
	}
	if args.ConfirmationToken != nil {
		user.ConfirmationToken = *args.ConfirmationToken
		user.ConfirmationTokenCreatedAt = time.Now()
	}

	user.UpdatedAt = time.Now()

	spew.Dump("user in service", user)
	if err = us.db.Update(user); err != nil {
		// TODO: better err messages (contstraints, ...)
		return nil, err
	}
	return user, nil
}

func (us *UserService) ChangePassword(ctx context.Context, userUUID uuid.UUID, newPassword string) (entity.User, error) {
	var err error
	user := entity.User{UUID: userUUID, Confirmed: true}
	if user.Password, err = hashPassword(newPassword); err != nil {
		return user, err
	}
	_, err = us.db.Model(&user).
		Set("password = ?password").
		Set("confirmation_token = ?", nil).
		Set("confirmation_token_created_at = ?", nil).
		Where("uuid = ?uuid").
		Returning("*").
		Update()

	return user, err
}

func (us *UserService) Confirm(ctx context.Context, tkn string) error {
	user := &entity.User{}
	err := us.db.Model(user).Where("confirmation_token = ?", tkn).First()
	if err != nil {
		return checkPgNotFoundErr(err)
	}
	ctx = context.WithValue(ctx, UserCtxKey, user)

	updateArgs := UserUpdateArgs{Confirmed: pointer.ToBool(true), ConfirmationToken: pointer.ToString("")}
	if _, err = us.Update(ctx, user.UUID, updateArgs); err != nil {
		return err
	}

	return nil
}

var UserCtxKey = &contextKey{"user"}

// AddUserToCtx does this:
// get token
// decode claims, get user uuid from token
// get user (passwd, lastLogin) from db (or cache)
// verify token
// add user to context
// TODO: refactor maybe
func AddUserToCtx(ctx context.Context, ja *auth.JwtAuth, us *UserService) (context.Context, error) {
	tknStr, err := TknFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	userUUID, err := auth.UserUUIDFromToken(tknStr)
	if err != nil {
		return nil, err
	}

	// TODO: add expiration check
	// TODO: use cache maybe
	user, err := us.GetByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	_, err = ja.Verify(ctx, tknStr, user)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't verify jwt token")
	}

	return context.WithValue(ctx, UserCtxKey, &user), nil
}

func UserFromCtx(ctx context.Context) *entity.User {
	user, _ := ctx.Value(UserCtxKey).(*entity.User)
	return user
}

// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}
