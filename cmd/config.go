package cmd

import (
	"fmt"
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
	case "local-docker":
		return localEnvironment(env)
	default:
		env = "local"
		return localEnvironment(env)
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

func localEnvironment(env string) *model.Config {
	path := fmt.Sprintf("./env/%s/", env)
	var appModel model.AppConfig
	appConfig := viper.New()
	appConfig.SetConfigName("app")
	appConfig.SetConfigType("yaml")
	appConfig.AddConfigPath(path)
	appConfig.ReadInConfig()
	appConfig.Unmarshal(&appModel)

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
		AppConfig: appModel,
		RdbConfig: rdbModel,
		JwtConfig: jwtModel,
	}
}

func uatEnvironment(env string) *model.Config {
	return &model.Config{}
}
