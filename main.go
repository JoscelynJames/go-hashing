package main

import (
	h "github.com/joscelynjames/go-hashing/handlers"
	s "github.com/joscelynjames/go-hashing/shutdown"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	// notify stopChan when a interrupt signal is recived
	signal.Notify(stopChan, os.Interrupt)
	// set up router/port/server
	mux := http.NewServeMux()
	port := ":8081"
	srv := &http.Server{Addr: port, Handler: mux}
	// must be in go routine in order to achieve the graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	// route for the post
	mux.HandleFunc("/", h.MakePosty)
	// send the data to stopChan
	<-stopChan // wait for server to get kill signal (SIGINT)
	// the function in package shutdown that gracefully shutdowns the server
	s.GracefullyShutdown(srv)
}
