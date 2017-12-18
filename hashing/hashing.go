package hashing

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// HashItOut function for MakePosty
func HashItOut(s string) string {
	if len(s) <= 1 {
		return "Please enter a password"
	}

	// create a new sha512
	sha := sha512.New()
	// encode the slice of bytes to a base64 string
	encoded := base64.URLEncoding.EncodeToString(sha.Sum(nil))
	fmt.Println("encoded", encoded)
	return encoded
}
