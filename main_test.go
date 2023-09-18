package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

// Test JWKS endpoint
func TestJWKSHandler(t *testing.T) {
	// Generate the RSA key pair for the test
	generateKeyPair()

	req, err := http.NewRequest("GET", "/.well-known/jwks.json", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(JWKSHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the returned key is valid
	var jwks map[string][]map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &jwks)
	if _, ok := jwks["keys"][0]["n"]; !ok {
		t.Errorf("handler returned unexpected body: missing key 'n'")
	}
}

// Test JWT generation
func TestGenerateJWTHandler(t *testing.T) {
	// Generate the RSA key pair for the test
	generateKeyPair()

	req, err := http.NewRequest("GET", "/generate-jwt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateJWTHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	tokenString := strings.TrimSpace(rr.Body.String())
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		t.Errorf("failed to parse generated JWT: %v", err)
	}
}

// TODO: Add more specific tests, error scenarios, etc.
