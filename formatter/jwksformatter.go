package jwksformatter

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

// JWK is a JSON Web Key
type JWK struct {
	Kid     string   `json:"kid"`
	Kty     string   `json:"kty"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	Use     string   `json:"use"`
	X5C     []string `json:"x5c"`
	X5T     string   `json:"x5t"`
	X5U     string   `json:"x5u"`
	X5TS256 string   `json:"x5t#S256"`
}

// JWKS key set which we will be using
type JWKS struct {
	Keys []JWK `json:"keys"`
}

func decodeX509string(cert string) (*x509.Certificate, error) {
	if !strings.Contains(cert, "CERTIFICATE-----") {
		cert = "-----BEGIN CERTIFICATE-----\n" + cert + "\n-----END CERTIFICATE-----\n"
	}
	block, _ := pem.Decode([]byte(cert))
	if block == nil {
		return nil, fmt.Errorf("invalid PEM block")
	}
	return x509.ParseCertificate(block.Bytes)
}

// Load a given keyset from a web source
func (jwks *JWKS) Load(uri string) error {

	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got non OK status code: %v", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(jwks); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}

// Get a given key
func (jwks *JWKS) Get(kid string) (JWK, error) {
	for _, k := range jwks.Keys {
		if k.Kid == kid {
			return k, nil
		}
	}

	return JWK{}, fmt.Errorf("key not found")
}

// GetAll keys
func (jwks *JWKS) GetAll() []JWK {
	return jwks.Keys
}

// Format a given JWKS to a format template
func (jwks *JWKS) Format(format string) (string, error) {
	t, err := template.New("jwks").Parse(format)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, jwks); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Expires returns the expiration date for the key
func (jwk *JWK) Expires(format string) string {
	cert, err := decodeX509string(jwk.X5C[0])
	if err != nil {
		return ""
	}
	return cert.NotAfter.Format(format)
}

// Issuer returns the issuing org name for the key
func (jwk *JWK) Issuer() string {
	cert, err := decodeX509string(jwk.X5C[0])
	if err != nil {
		return ""
	}
	return cert.Issuer.String()
}

// Subject returns the Subject for the key
func (jwk *JWK) Subject() string {
	cert, err := decodeX509string(jwk.X5C[0])
	if err != nil {
		return ""
	}
	return cert.Subject.String()
}

// Serial returns the Serial number for the key
func (jwk *JWK) Serial() string {
	cert, err := decodeX509string(jwk.X5C[0])
	if err != nil {
		return ""
	}
	return cert.SerialNumber.String()
}

// UUID returns a UUID for this key (this will be random and non-unique)
func (jwk *JWK) UUID() string {
	uuid, _ := uuid.NewRandom()
	return uuid.String()
}
