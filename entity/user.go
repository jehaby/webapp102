package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	UUID         uuid.UUID
	Name         string
	Email        string
	Password     string
	Role         UserRole
	DefaultPhone *uuid.UUID

	// https://github.com/go-pg/pg/issues/518
	// very inconvenient shit; chance to contribute!
	Confirmed         bool `sql:",notnull"`
	ConfirmationToken string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastLogout        time.Time
	BannedAt          time.Time
	BannedInfo        *string
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
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
