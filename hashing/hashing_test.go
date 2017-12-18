package hashing

import (
	"testing"
)

func TestHashItOut(t *testing.T) {
	
	test1 := HashItOut("password")
	if test1 != "sQnzu7wkTrgkQZF-0G1hi5AI3Qmzvv0bXgc5THBqi7mAsdd4Xll27ASbRt9fEyavWi6m0QP9B8lThf-rDKy8hg==" {
		t.Error("Password does not match sha512")
	}

	test2 := HashItOut(" ")
	if test2 != "Please enter a password" {
		t.Error("Did not ask for password")
	}
}
