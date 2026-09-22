package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/goccy/go-yaml"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/spf13/cobra"
)

var (
	claimsFilePath string
	keyFilePath    string
)

var gentokenCmd = &cobra.Command{
	Use:   "gentoken",
	Short: "Generate a JWT access token for testing or administration",
	Long: `Generates and signs a JSON Web Token (JWT). If a private key is not provided,
a new ECDSA P-256 key pair will be generated. Optional YAML claims file can be specified.`,
	RunE: func(c *cobra.Command, args []string) error {
		var key jwk.Key
		var err error

		if keyFilePath != "" {
			keyBytes, err := os.ReadFile(keyFilePath)
			if err != nil {
				return errl.Errorf("failed to read key file: %w", err)
			}
			key, err = jwk.ParseKey(keyBytes)
			if err != nil {
				return errl.Errorf("failed to parse JWK key: %w", err)
			}
		} else {
			privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				return errl.Errorf("failed to generate private key: %w", err)
			}
			key, err = jwk.FromRaw(privateKey)
			if err != nil {
				return errl.Errorf("failed to create JWK from raw key: %w", err)
			}
			_ = key.Set(jwk.KeyIDKey, "key-12345")
			_ = key.Set(jwk.KeyTypeKey, jwa.EC)
			_ = key.Set(jwk.AlgorithmKey, jwa.ES256)

			jwkJSON, _ := json.MarshalIndent(key, "", "  ")
			fmt.Fprintf(os.Stderr, "--- Generated Private Key (JWK) ---\n%s\n\n", string(jwkJSON))
		}

		token := jwt.New()

		if claimsFilePath != "" {
			claimsBytes, err := os.ReadFile(claimsFilePath)
			if err != nil {
				return errl.Errorf("failed to read claims file: %w", err)
			}
			var claims map[string]any
			if err := yaml.Unmarshal(claimsBytes, &claims); err != nil {
				return errl.Errorf("failed to parse YAML claims: %w", err)
			}

			if kid, ok := claims["kid"].(string); ok {
				_ = key.Set(jwk.KeyIDKey, kid)
				delete(claims, "kid")
			}

			jsonBytes, err := json.Marshal(claims)
			if err != nil {
				return errl.Errorf("failed to marshal claims to JSON: %w", err)
			}
			if err := json.Unmarshal(jsonBytes, token); err != nil {
				return errl.Errorf("failed to unmarshal claims into token: %w", err)
			}
		} else {
			_ = token.Set(jwt.IssuerKey, "https://issuer.example.com")
			_ = token.Set(jwt.AudienceKey, "https://verifier.example.com")
			_ = token.Set(jwt.SubjectKey, "did:key:12345")
			_ = token.Set("scope", "read write")
		}

		now := time.Now()
		_ = token.Set(jwt.IssuedAtKey, now)
		_ = token.Set(jwt.NotBeforeKey, now)
		_ = token.Set(jwt.ExpirationKey, now.AddDate(1, 0, 0))

		tokenBytes, err := jwt.Sign(token, jwt.WithKey(jwa.ES256, key))
		if err != nil {
			return errl.Errorf("failed to sign token: %w", err)
		}

		fmt.Println("--- Generated Token (JWT) ---")
		fmt.Println(string(tokenBytes))

		pubKey, err := key.PublicKey()
		if err == nil {
			_, err = jwt.Parse(tokenBytes, jwt.WithKey(jwa.ES256, pubKey))
			if err == nil {
				fmt.Fprintf(os.Stderr, "\nToken verified successfully!\n")
			}
		}

		return nil
	},
}

func init() {
	gentokenCmd.Flags().StringVar(&claimsFilePath, "claims", "", "Path to YAML file with claims")
	gentokenCmd.Flags().StringVar(&keyFilePath, "key", "", "Path to JWK file containing the private key")
	rootCmd.AddCommand(gentokenCmd)
}
