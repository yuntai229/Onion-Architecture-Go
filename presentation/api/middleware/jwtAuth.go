package middleware

import (
	"fmt"
	"onion-architecrure-go/domain/constant"
	"onion-architecrure-go/domain/model"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtAuthMiddleware struct {
	Config *model.Config
	Logger *zap.Logger
}

func NewJwtMiddleware(config *model.Config, logger *zap.Logger) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{config, logger}
}

func (middleware *JwtAuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := fmt.Sprintf("%v", ctx.Value("requestId"))

		tokenArr := ctx.Request.Header["Authorization"]
		if len(tokenArr) == 0 {
			newErr := constant.MissingTokenErr
			res := model.NewResModel().ResWithFail(newErr)
			middleware.Logger.Info("[ApiMiddleware][jwtAuth][Auth] Request end - Missing Token",
				zap.String("requestId", requestId),
				zap.Any("res", res),
			)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(tokenArr[0], "Bearer ")[1]
		middleware.Logger.Info("[ApiMiddleware][jwtAuth][Auth] Entry",
			zap.String("requestId", requestId),
			zap.String("tokenString", tokenString),
		)

		isValid, claims := model.NewJwtModel().VerifyJwt(middleware.Config.AppConfig.JwtKey, tokenString)
		if !isValid {
			newErr := constant.TokenInvalidErr
			res := model.NewResModel().ResWithFail(newErr)
			middleware.Logger.Info("[ApiMiddleware][jwtAuth][Auth] Request end - Token auth fail",
				zap.String("requestId", requestId),
				zap.Any("res", res),
			)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		ctx.Set("UserId", claims.UserID)
		ctx.Next()
	}
}
