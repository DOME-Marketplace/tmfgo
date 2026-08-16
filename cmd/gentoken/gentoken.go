package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func main() {
	var claimsFile string
	var keyFile string
	flag.StringVar(&claimsFile, "claims", "", "Path to YAML file with claims")
	flag.StringVar(&keyFile, "key", "", "Path to JWK file containing the private key")
	flag.Parse()

	var key jwk.Key
	var err error

	if keyFile != "" {
		// Load existing key
		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			log.Fatalf("Failed to read key file: %v", err)
		}
		key, err = jwk.ParseKey(keyBytes)
		if err != nil {
			log.Fatalf("Failed to parse JWK key: %v", err)
		}
	} else {
		// Generate a new private key (ECDSA P-256)
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			log.Fatalf("Failed to generate private key: %v", err)
		}
		key, err = jwk.FromRaw(privateKey)
		if err != nil {
			log.Fatalf("Failed to create JWK from raw key: %v", err)
		}
		// Set default Key ID and Algorithm if generating new
		key.Set(jwk.KeyIDKey, "key-12345")
		key.Set(jwk.KeyTypeKey, jwa.EC)
		key.Set(jwk.AlgorithmKey, jwa.ES256)

		// Print the generated key to screen (stderr) so it can be reused
		jwkJSON, _ := json.MarshalIndent(key, "", "  ")
		fmt.Fprintf(os.Stderr, "--- Generated Private Key (JWK) ---\n%s\n\n", string(jwkJSON))
	}

	// Create the Token (JWT)
	token := jwt.New()

	if claimsFile != "" {
		// Load claims from YAML
		claimsBytes, err := os.ReadFile(claimsFile)
		if err != nil {
			log.Fatalf("Failed to read claims file: %v", err)
		}
		var claims map[string]interface{}
		if err := yaml.Unmarshal(claimsBytes, &claims); err != nil {
			log.Fatalf("Failed to parse YAML claims: %v", err)
		}

		// Handle 'kid' from YAML if present
		if kid, ok := claims["kid"].(string); ok {
			key.Set(jwk.KeyIDKey, kid)
			delete(claims, "kid") // Don't put 'kid' in the payload
		}

		// Convert map to JSON and then unmarshal into jwt.Token to handle standard claims correctly
		jsonBytes, err := json.Marshal(claims)
		if err != nil {
			log.Fatalf("Failed to marshal claims to JSON: %v", err)
		}
		if err := json.Unmarshal(jsonBytes, token); err != nil {
			log.Fatalf("Failed to unmarshal claims into token: %v", err)
		}
	} else {
		// Default claims if no YAML provided
		token.Set(jwt.IssuerKey, "https://issuer.example.com")
		token.Set(jwt.AudienceKey, "https://verifier.example.com")
		token.Set(jwt.SubjectKey, "did:key:12345")
		token.Set("scope", "read write")
	}

	// Always set IssuedAt and Expiration (1 year) claims
	now := time.Now()
	token.Set(jwt.IssuedAtKey, now)
	token.Set(jwt.NotBeforeKey, now)
	token.Set(jwt.ExpirationKey, now.AddDate(1, 0, 0))

	// Sign the Token
	tokenBytes, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, key))
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	fmt.Println("--- Generated Token (JWT) ---")
	fmt.Println(string(tokenBytes))

	// Verification (Optional)
	pubKey, err := key.PublicKey()
	if err == nil {
		_, err = jwt.Parse(tokenBytes, jwt.WithKey(jwa.ES256, pubKey))
		if err == nil {
			fmt.Fprintf(os.Stderr, "\nToken verified successfully!\n")
		}
	}
}
