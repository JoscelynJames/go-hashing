package handlers

import (
	"fmt"
	"github.com/joscelynjames/go-hashing/hashing"
	"net/http"
	"time"
)

// MakePosty handles the 5 second sleep, retreives the password and hashes that password
func MakePosty(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside MakePosty")
	// keep the socket open but wait for 5 seconds
	time.Sleep(time.Second * 5)
	// parse the incomming data to get the value corrisponding to the key password
	r.ParseForm()
	x := r.Form.Get("password")
	fmt.Println("password incomming as", x)
	// pass the parsed value into the hashing function
	hashed := hashing.HashItOut(x)
	fmt.Fprintf(w, hashed)
}
