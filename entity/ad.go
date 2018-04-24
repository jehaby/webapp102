package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

type Ad struct {
	UUID        uuid.UUID `db:"uuid" sql:",pk" json:"uuid"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`

	UserUUID uuid.UUID `pg:"user_uuid"`
	User     *User     `pg:"fk:user_uuid" json:"user"`

	CategoryID int64 `db:"category_id"`
	Category   *Category

	ProductID int64
	Product   *Product

	LocalityID int64
	Locality   *Locality

	Price    int64
	Currency Currency

	BrandID int64
	Brand   *Brand

	Weight     int64
	Properties string
	CreatedAt  time.Time  `db:"created_at" sql:"default:now()"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time
}
