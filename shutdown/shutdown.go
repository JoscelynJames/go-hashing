package shutdown

import (
	"context"
	"log"
	"net/http"
)

// GracefullyShutdown shuts down the server when signal is received
func GracefullyShutdown(srv *http.Server) {
	log.Println("Shutting down server...")
	// get context and return the context with a new Done channel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// shutdown the sever
	srv.Shutdown(ctx)
	log.Println("gracefully stopped stopped")
}
