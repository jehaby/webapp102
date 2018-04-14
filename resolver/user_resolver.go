package resolver

import (
	"errors"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jehaby/webapp102/entity"
)

type userResolver struct {
	e *entity.User
}

func newUserResolver(e *entity.User) (*userResolver, error) {
	if e == nil {
		return nil, errors.New("newUserResolver: passed nil entity")
	}
	return &userResolver{e}, nil
}

func (ur *userResolver) Uuid() graphql.ID {
	return graphql.ID(ur.e.UUID.String())
}

func (ur *userResolver) Name() string {
	return ur.e.Name
}

func (ur *userResolver) Email() string {
	return ur.e.Email
}
