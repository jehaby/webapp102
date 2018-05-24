package config

// C is config
type C struct {
	Auth
	HTTP
	DB
}

type HTTP struct {
	Addr            string
	SecureJWTCookie bool
}

type Auth struct {
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
