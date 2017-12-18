package hashing

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func HashItOut(s string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(s))
	sha := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	fmt.Println(sha)
	return sha
}
