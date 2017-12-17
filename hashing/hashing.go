package hashing

import (
	"crypto/sha512"
	"io"
	"crypto/md5"
	"fmt"
	"encoding/base64"
)

func HashItOut(s string) {
	md5 := md5.New()
	io.WriteString(md5, s)
	sha_512 := sha512.New()
	sha_512.Write([]byte(s))
	sha := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))
	fmt.Println(sha)
}