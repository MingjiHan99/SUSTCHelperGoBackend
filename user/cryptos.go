package user

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5Encryption(str *string) string {
	h := md5.New()
	h.Write([]byte(*str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256Encryption(str *string) string {
	h := sha256.New()
	h.Write([]byte(*str))
	return hex.EncodeToString(h.Sum(nil))
}
