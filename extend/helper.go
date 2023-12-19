package extend

import (
	"crypto/md5"
	"encoding/hex"
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
	type authClaims struct {
		jwt.StandardClaims
		UserID uint
	}
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   "User",
			ExpiresAt: expiresAt,
		},
		UserID: Id,
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
