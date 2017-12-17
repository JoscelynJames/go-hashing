package hashing

import (
	"crypto/sha512"
	"io"
	"crypto/md5"
	"fmt"
)

func HashItOut(s string) {
	md5 := md5.New()
	io.WriteString(md5, s)
	sha_512 := sha512.New()
	sha_512.Write([]byte(s))
	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))
}