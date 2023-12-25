package extend

import (
	"crypto/md5"
	"encoding/hex"
)

var Helper HelperTools

type HelperTools struct{}

func (h *HelperTools) Hash(value string) string {
	hashBytes := md5.Sum([]byte(value))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}
