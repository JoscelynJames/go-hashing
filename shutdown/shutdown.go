package shutdown

import (
	"context"
	"log"
	"net/http"
)

func GracefullyShutdown(srv *http.Server) {
	log.Println("Shutting down server...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("gracefully stopped stopped")
}
