package main

import (
    "log"
    "os"
    "strconv"
)

var (
    jwksRetention int
    jwtIssuer     string
    jwtSubject    string
    jwtAudience   string
    jwtTTL        int
    jwksKeyTTL    int
) 

// init function is automatically executed before the main function and is typically used for initialization.
func init() {
    // Read environment variables using helper functions to retrieve and validate their values.
    jwksRetention := getEnvInt("JWKS_RETENTION", 5) // Number of keys to retain.
    jwtIssuer := getEnv("JWT_ISSUER", "", true) // Issuer of the JWT.
    jwtSubject := getEnv("JWT_SUBJECT", "", true) // Subject of the JWT.
    jwtAudience := getEnv("JWT_AUDIENCE", "", true) // Audience of the JWT.
    jwtTTL := getEnvInt("JWT_TTL", 60) // Time-To-Live for the JWT.
    jwksKeyTTL := getEnvInt("JWKS_KEY_TTL", 60) // Time-To-Live for the JWKS key.
    
    // Calculate maximum safe duration a JWT can be valid (based on the retention of keys and the key's TTL).
    // This ensures that a JWT doesn't outlive the key used to sign it.
    maxSafeJWTTTL := jwksRetention * jwksKeyTTL
    // Validate that the JWT's TTL doesn't exceed the safe maximum.
    if jwtTTL > maxSafeJWTTTL {
        log.Fatalf("JWT_TTL (%d) exceeds the maximum safe value (%d)", jwtTTL, maxSafeJWTTTL)
    } 
}

// getEnv is a helper function that reads a string value from environment variables.
// If the environment variable is essential and not set, the application exits.
// Otherwise, it returns the default value.
func getEnv(key, defaultValue string, essential bool) string {
    value := os.Getenv(key)
    if value == "" {
        if essential {
            log.Fatalf("Essential environment variable %s not set", key)
        }
        return defaultValue
    }
    return value
}   

// getEnvInt is a helper function that reads an integer value from environment variables.
// It performs the conversion from string to integer and handles errors. If the conversion fails,
// the application exits with an error message.
func getEnvInt(key string, defaultValue int) int {
    valueStr := os.Getenv(key)
    if valueStr == "" {
    	return defaultValue
    }
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        log.Fatalf("Invalid value for %s: %v", key, err)
    }
    return value
} 