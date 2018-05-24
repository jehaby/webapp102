package service

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/jehaby/webapp102/entity"
	"github.com/jehaby/webapp102/service/auth"
	"github.com/jehaby/webapp102/storage"
)

type UserService struct {
	Repo *storage.UserRepository // TODO: make unexported maybe
}

func newUserService(db *sqlx.DB) *UserService {
	return &UserService{
		Repo: storage.NewUserRepository(db),
	}
}

var UserCtxKey = &contextKey{"user"}

func AddUserToCtx(ctx context.Context, ja *auth.JwtAuth, us *UserService) (context.Context, error) {
	tkn, err := ja.Verify(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't verify jwt token")
	}

	// TODO: improve this
	userUUID, _ := auth.Claims(tkn.Claims.(jwt.MapClaims))["user"].(map[string]interface{})["UUID"].(string)
	user, err := us.Repo.GetByUUID(uuid.FromStringOrNil(userUUID))
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, UserCtxKey, user), nil
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
