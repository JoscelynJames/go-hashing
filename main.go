package main 

import (
	"net/http"

	h "github.com/joscelynjames/go-hashing/handlers"
)

func main() {

	http.HandleFunc("/makePosty", h.MakePosty)
	http.ListenAndServe(":8081", nil)
}