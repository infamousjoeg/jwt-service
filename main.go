package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func generateKeyPair() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating RSA keys: %v", err)
	}
	publicKey = &privateKey.PublicKey
	log.Println("Generated new RSA key pair.")
}

// JWKSHandler serves the JWKS endpoint
func JWKSHandler(w http.ResponseWriter, r *http.Request) {
	jwks := map[string]interface{}{
		"keys": []interface{}{
			map[string]interface{}{
				"kty": "RSA",
				"alg": "RS256",
				"use": "sig",
				"n":   publicKey.N,
				"e":   publicKey.E,
			},
		},
	}
	json.NewEncoder(w).Encode(jwks)
}

// GenerateJWTHandler generates a JWT
func GenerateJWTHandler(w http.ResponseWriter, r *http.Request) {
	ttlStr := os.Getenv("JWT_TTL")
	if ttlStr == "" {
		ttlStr = "60" // default to 60 minutes if not provided
	}
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		log.Fatalf("Invalid JWT_TTL value: %v", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": os.Getenv("JWT_ISSUER"),
		"sub": os.Getenv("JWT_SUBJECT"),
		"aud": os.Getenv("JWT_AUDIENCE"),
		"exp": time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(tokenString))
}

func main() {
	// Initial key generation
	generateKeyPair()

	// Set up periodic key rotation based on TTL from environment variable
	ttlStr := os.Getenv("JWKS_KEY_TTL")
	if ttlStr == "" {
		ttlStr = "60" // default to 60 minutes if not provided
	}
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		log.Fatalf("Invalid JWKS_KEY_TTL value: %v", err)
	}

	go func() {
		for {
			time.Sleep(time.Duration(ttl) * time.Minute)
			generateKeyPair()
		}
	}()

	http.HandleFunc("/.well-known/jwks.json", JWKSHandler)
	http.HandleFunc("/generate-jwt", GenerateJWTHandler)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
