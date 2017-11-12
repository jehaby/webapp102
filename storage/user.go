package storage

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jehaby/webapp102/entity"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

type UserRepository struct {
	db *sqlx.DB
}

func (ur *UserRepository) GetByName(name string) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.Get(user, "SELECT * FROM users WHERE name=$1", name)
	if err != nil {
		return nil, errors.Wrapf(err, "getting user %+v")
	}

	return user, nil
}

func (ur *UserRepository) Save(u entity.User) error {

	if u.UUID == uuid.Nil {
		u.UUID = uuid.NewV4()
	}

	spew.Dump(u)

	res, err := ur.db.Exec(`INSERT INTO  users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`,
		u.UUID,
		u.Name,
		u.Email,
		u.Password,
		time.Now(),
	)

	if err != nil {
		return errors.Wrapf(err, "saving user %+v", u)
	}
	if ra, err := res.RowsAffected(); err != nil || ra != 1 {
		return errors.Wrapf(err, "saving user %+v, rows affected: %d", u, ra)
	}

	return nil

}
