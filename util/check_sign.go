package util

import (
	"crypto/hmac"
	"crypto/sha256"
)

func CheckHMAC(msg, msgMac, key []byte) bool {
	mac := hmac.New(sha256.New,key)
	mac.Write(msg)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(expectedMAC, msgMac)
}