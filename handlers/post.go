package handlers

import (
	"fmt"
	"github.com/joscelynjames/go-hashing/hashing"
	"net/http"
	"time"
)

func MakePosty(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println("Inside MakePosty")

	time.Sleep(time.Second * 5)

	fmt.Println("Go routine finished")
	fmt.Println("Hashing started")

	r.ParseForm()
	x := r.Form.Get("password")
	fmt.Println("password incomming as", x)

	hashed := hashing.HashItOut(x)
	fmt.Fprintf(w, hashed)
}
