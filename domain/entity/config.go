package entity

type Config struct {
	Env       string
	RdbConfig RdbConfig
	JwtConfig JwtConfig
}

type RdbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Db       string
	Charset  string
	Timeout  string
}

type JwtConfig struct {
	Key string
}
