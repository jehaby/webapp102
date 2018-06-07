package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	UUID         uuid.UUID  `json:"uuid"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"-"`
	Role         UserRole   `json:"role"`
	DefaultPhone *uuid.UUID `json:"default_phone,omitempty"`

	Phones []*Phone `json:"phones,omitempty"`

	// https://github.com/go-pg/pg/issues/518
	// very inconvenient shit; chance to contribute!
	Confirmed                  bool      `sql:",notnull" json:"confirmed"`
	ConfirmationToken          string    `json:"-"`
	ConfirmationTokenCreatedAt time.Time `json:"-"`
	CreatedAt                  time.Time `json:"-"`
	UpdatedAt                  time.Time `json:"-"`
	LastLogout                 time.Time `json:"-"`
	BannedAt                   time.Time `json:"-"`
	BannedInfo                 *string   `json:"-"`
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
