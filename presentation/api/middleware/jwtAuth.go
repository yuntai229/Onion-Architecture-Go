package middleware

import (
	domain "onion-architecrure-go/domain/entity"
	"onion-architecrure-go/extend"
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
			newErr := domain.MissingTokenErr
			res := domain.Response.ResWithFail(newErr)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(tokenArr[0], "Bearer ")[1]

		isValid, claims := extend.Helper.VerifyJwt(tokenString)
		if !isValid {
			newErr := domain.TokenInvalidErr
			res := domain.Response.ResWithFail(newErr)
			ctx.JSON(newErr.HttpCode, res)
			ctx.Abort()
			return
		}

		ctx.Set("UserId", claims.UserID)
		ctx.Next()
	}
}
