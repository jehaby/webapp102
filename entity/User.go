package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	UUID      uuid.UUID  `db:"uuid"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	LastLogin *time.Time `db:"last_login"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
