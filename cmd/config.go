package cmd

import (
	"onion-architecrure-go/domain/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitAppEnv() *model.Config {
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

func defaultEnvironment(env string) *model.Config {
	path := "./env/local/"

	var rdbModel model.RdbConfig
	rdbConfig := viper.New()
	rdbConfig.SetConfigName("rdb")
	rdbConfig.SetConfigType("yaml")
	rdbConfig.AddConfigPath(path)
	rdbConfig.ReadInConfig()
	rdbConfig.Unmarshal(&rdbModel)

	var jwtModel model.JwtConfig
	jwtConfig := viper.New()
	jwtConfig.SetConfigName("jwt")
	jwtConfig.SetConfigType("yaml")
	jwtConfig.AddConfigPath(path)
	jwtConfig.ReadInConfig()
	jwtConfig.Unmarshal(&jwtModel)

	return &model.Config{
		Env:       env,
		RdbConfig: rdbModel,
		JwtConfig: jwtModel,
	}
}

func uatEnvironment(env string) *model.Config {
	return &model.Config{}
}
