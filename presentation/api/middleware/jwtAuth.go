package middleware

import (
	"fmt"
	"onion-architecrure-go/domain/entity"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtAuthMiddleware struct {
	Config *entity.Config
	Logger *zap.Logger
}

func NewJwtMiddleware(config *entity.Config, logger *zap.Logger) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{config, logger}
}

func (middleware *JwtAuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := fmt.Sprintf("%v", ctx.Value("requestId"))

		tokenArr := ctx.Request.Header["Authorization"]
		if len(tokenArr) == 0 {
			newErr := entity.MissingTokenErr
			res := entity.Response.ResWithFail(newErr)
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

		var authClaims entity.UserAuthClaims
		isValid, claims := authClaims.VerifyJwt(middleware.Config.JwtConfig.Key, tokenString)
		if !isValid {
			newErr := entity.TokenInvalidErr
			res := entity.Response.ResWithFail(newErr)
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
