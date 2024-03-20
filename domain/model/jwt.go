package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserAuthClaims struct {
	jwt.StandardClaims
	UserID uint
}

func NewJwtModel() *UserAuthClaims {
	return &UserAuthClaims{}
}

func (model *UserAuthClaims) VerifyJwt(key, tokenString string) (bool, *UserAuthClaims) {
	token, err := jwt.ParseWithClaims(tokenString, model, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil || !token.Valid {
		return false, nil
	}
	return true, model
}

func (model *UserAuthClaims) GenJwt(key string) (string, error) {
	var jwtKey = []byte(key)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, UserAuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "User",
			ExpiresAt: expiresAt,
		},
		UserID: model.UserID,
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
