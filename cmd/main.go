package main

import (
	"log"
	"net/http"

	"github.com/shirocola/NBake-BE/internal/auth"
)

func main() {
	// Setup HTTP server
	mux := http.NewServeMux()

	// Setup routes using your auth package
	// Assuming SetupRoutes returns an http.Handler with all the auth routes configured
	authHandler := auth.SetupRoutes()
	mux.Handle("/auth/", http.StripPrefix("/auth", authHandler))

	// Set up a server listening on port 8080
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	// Start the server
	log.Println("Starting server on :3000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
