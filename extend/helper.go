package extend

import (
	"crypto/md5"
	"encoding/hex"
	domain "onion-architecrure-go/domain/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

var Helper HelperTools

type HelperTools struct{}

func (h *HelperTools) Hash(value string) string {
	hashBytes := md5.Sum([]byte(value))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}

func (h *HelperTools) GenJwt(Id uint) (string, error) {
	var jwtKey = []byte("Test")
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, domain.UserAuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "User",
			ExpiresAt: expiresAt,
		},
		UserID: Id,
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func (h *HelperTools) VerifyJwt(tokenString string) (bool, *domain.UserAuthClaims) {
	var claims domain.UserAuthClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("Test"), nil
	})

	if err != nil || !token.Valid {
		return false, nil
	}
	return true, &claims
}
