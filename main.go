package main 

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	h "github.com/joscelynjames/go-hashing/handlers"
	s "github.com/joscelynjames/go-hashing/shutdown"
)

func main() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	
	mux := http.NewServeMux()
	port := ":8081"
	srv := &http.Server{Addr: port, Handler: mux}

	go func() {
		// service connections
		log.Printf("Running on port %v", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	
	mux.HandleFunc("/", h.MakePosty)

	<-stopChan // wait for server to get kill signal (SIGINT)
	
	s.GracefullyShutdown(srv)
}