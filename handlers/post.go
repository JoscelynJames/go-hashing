package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func waitFive(done chan bool) {
	fmt.Println("Inside waitFive")
	time.Sleep(time.Second * 5)
	done <- true
}

func MakePosty(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside MakePosty")

	done := make(chan bool)
	go waitFive(done)
	<-done

	fmt.Println("Go routine finished")
}


