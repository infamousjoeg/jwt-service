package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"os/context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// main is the primary entry point for the application.
func main() {
	// Generate the initial RSA key pair for JWT signing.
	generateKeyPair()

	// Periodically generate a new RSA key pair based on the TTL.
	go func() {
		for {
			time.Sleep(time.Duration(jwtTTL) * time.Minute)
			generateKeyPair()
		}
	}()

	// Register HTTP handlers for the JWKS and JWT generation endpoints.
	http.HandleFunc("/.well-known/jwks.json", JWKSHandler)
	http.HandleFunc("/generate-jwt", GenerateJWTHandler)

	// Configure and start the HTTP server.
	srv := &http.Server{
	    Addr: ":8080",
	    ReadTimeout: 5 * time.Second,
	    WriteTimeout: 10 * time.Second,
	    IdleTimeout: 15 * time.Second,
	    Handler: nil,
    }
    
    // Start server in a goroutine so it doesn't block the main thread.
    // This allows for graceful shutdown later on.
    go func() {
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("Server ListenAndServe(): %v", err)
        }
    }()
    
    // Setup graceful shutdown: Wait for an interrupt signal, then shutdown the HTTP server.
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <-stop
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server Shutdown Failed:%+v", err)
    }
    log.Println("Server gracefully stopped")    
}
