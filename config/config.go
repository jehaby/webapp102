package config

// C is config
type C struct {
	HTTP
	DB
}

type HTTP struct {
	Addr string
}

type DB struct {
	Conn string
}