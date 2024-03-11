package cmd

import (
	"onion-architecrure-go/domain/entity"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitAppEnv() *entity.Config {
	switch env := os.Getenv("APP_ENV"); env {
	// case "uat":
	// 	return uatEnvironment(env)
	default:
		env = "local"
		return defaultEnvironment(env)
	}
}

func InitLoggerEnv() zap.Config {
	switch env := os.Getenv("APP_ENV"); env {
	case "uat":
		return zap.NewProductionConfig()
	default:
		return zap.NewDevelopmentConfig()
	}
}

func InitRunningMode() {
	switch env := os.Getenv("APP_ENV"); env {
	case "uat":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func defaultEnvironment(env string) *entity.Config {
	path := "./env/local/"

	var rdbEntity entity.RdbConfig
	rdbConfig := viper.New()
	rdbConfig.SetConfigName("rdb")
	rdbConfig.SetConfigType("yaml")
	rdbConfig.AddConfigPath(path)
	rdbConfig.ReadInConfig()
	rdbConfig.Unmarshal(&rdbEntity)

	var jwtEntity entity.JwtConfig
	jwtConfig := viper.New()
	jwtConfig.SetConfigName("jwt")
	jwtConfig.SetConfigType("yaml")
	jwtConfig.AddConfigPath(path)
	jwtConfig.ReadInConfig()
	jwtConfig.Unmarshal(&jwtEntity)

	return &entity.Config{
		Env:       env,
		RdbConfig: rdbEntity,
		JwtConfig: jwtEntity,
	}
}

func uatEnvironment(env string) *entity.Config {
	return &entity.Config{}
}
