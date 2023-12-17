package cmd

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlSetting struct {
	User     string `mapstructure: user`
	Password string `mapstructure: password`
	Host     string `mapstructure: host`
	Port     string `mapstructure: port`
	Db       string `mapstructure: db`
	Charset  string `mapstructure: charset`
	Timeout  string `mapstructure: timeout`
}

var Db *gorm.DB

func InitDb() *gorm.DB {
	var option = "%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&loc=Local&parseTime=True"

	dsn := fmt.Sprintf(option,
		"root",
		"",
		"localhost",
		"3306",
		"gin",
		"utf8mb4",
		"12s",
	)

	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return Db
}
