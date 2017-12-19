package hashing

import (
	"testing"
)

func TestHashItOut(t *testing.T) {

	test1 := HashItOut("password")
	if test1 != "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg_SpIdNs6c5H0NE8XYXysP-DGNKHfuwvY7kxvUdBeoGlODJ6-SfaPg==" {
		t.Error("Password does not match sha512")
	}

	test2 := HashItOut("")
	if test2 != "Please enter a password" {
		t.Error("Did not ask for a password")
	}
}
