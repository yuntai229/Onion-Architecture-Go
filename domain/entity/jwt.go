package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserAuthClaims struct {
	jwt.StandardClaims
	UserID uint
}

func (entity *UserAuthClaims) VerifyJwt(tokenString string) (bool, *UserAuthClaims) {
	token, err := jwt.ParseWithClaims(tokenString, entity, func(token *jwt.Token) (interface{}, error) {
		return []byte("Test"), nil
	})

	if err != nil || !token.Valid {
		return false, nil
	}
	return true, entity
}

func (entity *UserAuthClaims) GenJwt() (string, error) {
	var jwtKey = []byte("Test")
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, UserAuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "User",
			ExpiresAt: expiresAt,
		},
		UserID: entity.UserID,
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
