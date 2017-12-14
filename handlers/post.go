package handlers

import (
	"net/http"
	"fmt"
)

func MakePosty(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside make posty")
}