package config

// C is config
type C struct {
	HTTP
	DB
}

type HTTP struct {
	Addr   string
	Secret string
}

type DB struct {
	Conn string

	User     string
	Database string
	Host     string
	Port     string
	SSLMode  bool
}
