package domain

import "github.com/golang-jwt/jwt"

type UserAuthClaims struct {
	jwt.StandardClaims
	UserID uint
}
