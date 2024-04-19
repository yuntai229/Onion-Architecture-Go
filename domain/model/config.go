package model

type Config struct {
	Env       string
	AppConfig AppConfig
	RdbConfig RdbConfig
	JwtConfig JwtConfig
}

type AppConfig struct {
	Version string
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
