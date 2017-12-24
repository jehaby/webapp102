package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

type Ad struct {
	UUID        uuid.UUID  `db:"uuid" json:"uuid"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	UserUUID    string     `db:"user_uuid"`
	User        *User      `json:"user"`
	CategoryID  uint16     `db:"category_id"`
	Component   *Component `json:"component"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}
