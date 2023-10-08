package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
    
    "github.com/golang-jwt/jwt"
) 

// JWKSHandler serves the JWKS (JSON Web Key Set) endpoint.
// It provides the RSA public keys in the keyRing in JWKS format.
func JWKSHandler(w http.ResponseWriter, r *http.Request) {
	// Slice to store the JWKS keys.
	var jwksKeys []map[string]interface{}
	
	// Iterate through each entry in the keyRing.
	for _, entry := range keyRing {
		// Construct the JWKS representation for the public key.
		keyData := map[string]interface{
			"kty": "RSA", // Key Type: RSA
			"alg": "RS256", // Algorithm: RS256
			"use": "sig", // Key Use: Signature
			"n":   entry.PublicKey.N, // RSA Public Key Modulus
			"e":   entry.PublicKey.E, // RSA Public Key Exponent
			"kid": entry.KID, // Key ID
		}
		// Append the key data to the jwksKeys slice.
		jwksKeys = append(jwksKeys, keyData)	
	}
	
	// Create the JWKS response object.
	jwks := map[string]interface{}{"keys": jwksKeys}
	// Send the JWKS response as JSON.
	json.NewEncoder(w).Encode(jwks)
}

// GenerateJWTHandler handles the JWT generation request.
// It creates a new JWT and signs it using the latest private key in the keyRing.
func GenerateJWTHandler(w http.ResponseWriter, r *http.Request) {
	// Use the helper function to convert the JWT_TTL string to integer and determine the token's validity duration.
	ttl := getEnvInt("JWT_TTL", 60)
	
	// Create a new JWT with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": os.Getenv("JWT_ISSUER"), // Issuer
		"sub": os.Getenv("JWT_SUBJECT"), // Subject
		"aud": os.Getenv("JWT_AUDIENCE"), // Audience
		"exp": time.Now().Add(time.Duration(ttl) * time.Minute).Unix(), // Expiry Time
	})
	
	// Set the 'kid' header to the ID of the latest key in the keyRing.
	token.Header["kid"] = keyRing[len(keyRing)-1].KID
	// Sign the token using the private key.
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		// Return an internal server error if there's an issue signing the token.
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	
	// Send the signed JWT as the response.
	w.Write([]byte(tokenString))
}