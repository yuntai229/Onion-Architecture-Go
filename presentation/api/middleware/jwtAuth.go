package middleware

import (
	"onion-architecrure-go/domain/entity"
	"strings"

	"github.com/gin-gonic/gin"
)

type JwtAuthMiddleware struct{}

func NewJwtMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (middleware *JwtAuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenArr := ctx.Request.Header["Authorization"]
		if len(tokenArr) == 0 {
			newErr := entity.MissingTokenErr
			res := entity.Response.ResWithFail(newErr)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(tokenArr[0], "Bearer ")[1]

		var authClaims entity.UserAuthClaims
		isValid, claims := authClaims.VerifyJwt(tokenString)
		if !isValid {
			newErr := entity.TokenInvalidErr
			res := entity.Response.ResWithFail(newErr)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		ctx.Set("UserId", claims.UserID)
		ctx.Next()
	}
}
