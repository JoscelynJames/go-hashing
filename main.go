package main 

import (
	"net/http"
	"context"
	"log"
	"os"
	"os/signal"
	// "time"
	h "github.com/joscelynjames/go-hashing/handlers"
	// s "github.com/joscelynjames/go-hashing/shutdown"
)

func main() {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	
	mux := http.NewServeMux()
	port := ":8081"

	mux.HandleFunc("/", h.MakePosty)

	srv := &http.Server{Addr: port, Handler: mux}

	go func() {
		// service connections
		log.Printf("Running on port %v", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan // wait for server to get kill signal (SIGINT)
	log.Println("Shutting down server...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	srv.Shutdown(ctx)

	log.Println("gracefully stopped stopped")

	http.ListenAndServe(":8081", nil)
}