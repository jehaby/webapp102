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
	Alg    string
	Secret string
}

type DB struct {
	Conn string
	URL  string

	User     string
	Database string
	Host     string
	Port     string
	SSLMode  bool
}
