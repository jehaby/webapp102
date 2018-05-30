package service

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service/auth"
	"github.com/jehaby/webapp102/storage"
)

type UserService struct {
	Repo *storage.UserRepository // TODO: make unexported maybe
	db   *pg.DB
}

func newUserService(db *sqlx.DB, pgdb *pg.DB) *UserService {
	return &UserService{
		Repo: storage.NewUserRepository(db),
		db:   pgdb,
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
