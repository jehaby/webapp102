package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	UUID         uuid.UUID  `db:"uuid"`
	Name         string     `db:"name"`
	Email        string     `db:"email"`
	Password     string     `db:"password"`
	Role         UserRole   `db:"role"`
	DefaultPhone *uuid.UUID `db:"default_phone"`
	LastLogin    *time.Time `db:"last_login"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
}

type UserRole uint8

const (
	RoleUser UserRole = 1 << iota
	RoleModerator
	RoleAdmin
)

func (r *UserRole) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("UserRole: Scan source is not []byte")
	}
	switch string(asBytes) {
	case "user":
		*r = RoleUser
	case "moderator":
		*r = RoleModerator
	case "admin":
		*r = RoleAdmin
	default:
		return fmt.Errorf("UserRole: Scan: unknown value: %v", asBytes)
	}
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	switch r {
	case RoleUser:
		return "user", nil
	case RoleModerator:
		return "moderator", nil
	case RoleAdmin:
		return "admin", nil
	}
	return nil, fmt.Errorf("unknown role")
}
