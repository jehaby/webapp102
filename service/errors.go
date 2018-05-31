package service

import (
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

var (
	ErrNotAuthorized = errors.New("not authorized")
	ErrNotFound      = errors.New("not found")
	ErrNotAllowed    = errors.New("not allowed")
)

func checkPgNotFoundErr(err error) error {
	if err == pg.ErrNoRows {
		return ErrNotFound
	}
	return err
}
