package model

type Config struct {
	Env       string
	AppConfig AppConfig
	RdbConfig RdbConfig
}

type AppConfig struct {
	Version string
	JwtKey  string
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
