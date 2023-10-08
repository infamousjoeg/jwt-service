package main

import (
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "log"
    
    "github.com/golang-jwt/jwt"
) 

// KeyEntry struct holds the RSA public key and the associated key ID.
// This structure is used to store the generated public keys and their respective IDs in the keyRing.
type KeyEntry struct {
    PublicKey *rsa.PublicKey
    KID       string
}

// Package-level variables.
var (
    keyRing       []KeyEntry // Slice that stores the public keys and their associated IDs.
    privateKey    *rsa.PrivateKey // RSA private key used for signing JWTs.
)

// computeKeyID takes an RSA public key and computes a unique key ID for it by hashing the key.
// The hash is then encoded in base64 to generate the key ID.
func computeKeyID(pub *rsa.PublicKey) string {
	pubASN1, err := x509.MarshalPKCS1PublicKey(pub)
	if err != nil {
		log.Fatalf("Error encoding RSA public key: %v", err)
	}

    // Create a new SHA-256 hasher.
	hasher := sha256.New()
	// Write the ASN1 encoding of the public key to the hasher.
	hasher.Write(pubASN1)
	// Compute the hash and encode it using base64 to form the key ID.
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// generateKeyPair is responsible for generating a new RSA key pair.
// It also computes the key ID for the public key and appends it to the keyRing.
// If the keyRing exceeds the jwksRetention count, the oldest key is removed.
func generateKeyPair() {
	var err error
	// Generate a new RSA key pair with a bit size of 2048.
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating RSA keys: %v", err)
	}
	// Compute the key ID for the newly generated public key.
	kid := computeKeyID(&privateKey.PublicKey)
	// Add the public key and its key ID to the keyRing.
	keyRing = append(keyRing, KeyEntry{PublicKey: &privateKey.PublicKey, KID: kid})
	log.Println("Generated new RSA key pair with kid:", kid)
	
	// Trim the keyRing to ensure it doesn't exceed the jwksRetention count.
	// If it does, remove the oldest keys.
	for len(keyRing) > jwksRetention {
	    keyRing = keyRing[1:]
    }  
}