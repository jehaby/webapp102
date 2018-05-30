package service

import (
	"context"
	"time"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/pkg/log"
	"github.com/jehaby/webapp102/service/auth"
	"github.com/jehaby/webapp102/storage"
)

type UserService struct {
	Repo *storage.UserRepository // TODO: make unexported maybe
	db   *pg.DB
	val  *validator.Validate
	log  *log.Logger
}

func newUserService(
	db *sqlx.DB,
	pgdb *pg.DB,
	val *validator.Validate,
	log *log.Logger,
) *UserService {
	return &UserService{
		Repo: storage.NewUserRepository(db),
		db:   pgdb,
		val:  val,
		log:  log,
	}
}

func (us *UserService) GetByNameOrEmail(nameOrEmail string) (entity.User, error) {
	res := entity.User{}
	err := us.db.Model(&res).Where("name = ?", nameOrEmail).WhereOr("email = ?", nameOrEmail).First()
	if err != nil {
		return res, errors.Wrapf(err, "couldn't get user by name or email (%s)", nameOrEmail)
	}
	return res, nil
}

func (us *UserService) GetByUUID(uuid uuid.UUID) (entity.User, error) {
	res := entity.User{}
	err := us.db.Model(&res).Where("uuid = ?", uuid).First()
	if err != nil {
		return res, errors.Wrapf(err, "couldn't get user by uuid (%s)", uuid)
	}
	return res, nil
}

func (us *UserService) Create(ctx context.Context) (entity.User, error) {
	res := entity.User{}
	// TODO: implement
	return res, nil
}

type UserUpdateArgs struct {
	Email      *string `validate:"omitempty,email"`
	LastLogout *time.Time
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
			// TODO: better error (404 and 500)
			return nil, err
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
	user.UpdatedAt = time.Now()
	if err = us.db.Update(user); err != nil {
		return nil, err
	}
	return user, nil
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
	tknStr, err := auth.TknFromCtx(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "no token in context")
	}

	userUUID, err := auth.UserUUIDFromToken(tknStr)
	if err != nil {
		return nil, errors.Wrapf(err, "getting user uuid from token (%s)", tknStr)
	}

	// TODO: add expiration check
	// TODO: use cache maybe
	user, err := us.GetByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	_, err = ja.Verify(ctx, user)
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
