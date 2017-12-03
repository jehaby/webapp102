package entity

import "github.com/lib/pq"

type Category struct {
	ID   int64         `db:"id" json:"-"`
	Name string        `db:"name" json:"name"`
	Path pq.Int64Array `db:"path" json:"path"` // TODO: uint16 array, probably PR in postgres driver
}
