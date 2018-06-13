package service

import (
	"github.com/davecgh/go-spew/spew"
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

func ignorePgNotFoundErr(err error) error {
	spew.Dump("err2", err)
	if err == nil {
		// TODO: might be nil interface bug
		return nil
	}
	if err == pg.ErrNoRows || err == ErrNotFound {
		return nil
	}
	return errors.Wrap(err, "bad pg error")
}

type ErrBadRequest struct {
	err error
}

func (e *ErrBadRequest) Error() string {
	return e.err.Error()
}

type IntegrityViolationer interface {
	error
	IntegrityViolation() bool
}
