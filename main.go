package main 

import (
	"net/http"

	h "github.com/joscelynjames/go-hashing/handlers"
)

func main() {

	http.HandleFunc("/", h.MakePosty)
	http.ListenAndServe(":8081", nil)
}