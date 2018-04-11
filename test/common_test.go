package test

import (
	"github.com/jmoiron/sqlx"
)

const (
	defaultPostgresConn = ""
	defaultHTTPAddr     = "http://localhost:8899/"

	jsonContentType = "application/json"
)

func db() *sqlx.DB {
	panic("not implemented")
	return nil
}

func httpAddr() string {
	return defaultHTTPAddr
}

func graphqlAddr() string {
	return httpAddr() + "query"
}
