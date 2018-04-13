package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

type Ad struct {
	UUID        uuid.UUID `db:"uuid" sql:",pk" json:"uuid"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	UserUUID    string    `pg:"user_uuid"`
	User        *User     `pg:"fk:user_uuid" json:"user"`
	CategoryID  uint16    `db:"category_id"`
	// Component   *Component // `pg:"fk:component_id" json:"component"`
	CreatedAt time.Time  `db:"created_at" sql:"default:now()" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
