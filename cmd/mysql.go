package cmd

import (
	"fmt"
	"onion-architecrure-go/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(config *entity.Config) *gorm.DB {
	var option = "%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&loc=Local&parseTime=True"

	dsn := fmt.Sprintf(option,
		config.RdbConfig.User,
		config.RdbConfig.Password,
		config.RdbConfig.Host,
		config.RdbConfig.Port,
		config.RdbConfig.Db,
		config.RdbConfig.Charset,
		config.RdbConfig.Timeout,
	)

	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return Db
}
