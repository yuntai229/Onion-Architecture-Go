package extend

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

var Helper HelperTools

type HelperTools struct{}

func (h *HelperTools) Hash(value string) string {
	hashBytes := md5.Sum([]byte(value))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}

func (h *HelperTools) FormatToTimeString(inputTime time.Time) string {
	return inputTime.Format("2006-01-02 15:04:05")
}
