package utils

import (
	"crypto/md5"
	"fmt"
)

func GetHash(input string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return ""
	}
	bytes := hash.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}
