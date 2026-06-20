package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

// Precomputed base64 RawURLEncoding for static JWT header: {"alg":"HS256","typ":"JWT"}
const jwtHeader = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

type TokenClaims struct {
	Admin bool  `json:"admin"`
	Exp   int64 `json:"exp"`
}

func getSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_key_change_me"
	}
	return []byte(secret)
}

func sign(message string) string {
	h := hmac.New(sha256.New, getSecret())
	h.Write([]byte(message))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// GenerateToken creates a signed HS256 JWT declaring admin authorization
func GenerateToken(duration time.Duration) (string, error) {
	claims := TokenClaims{
		Admin: true,
		Exp:   time.Now().Add(duration).Unix(),
	}
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.RawURLEncoding.EncodeToString(claimsBytes)
	signingInput := jwtHeader + "." + payload

	return signingInput + "." + sign(signingInput), nil
}

// VerifyToken decodes and validates the HS256 JWT, returning the claims if valid
func VerifyToken(tokenString string) (*TokenClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 || parts[0] != jwtHeader {
		return nil, errors.New("invalid token format")
	}

	signingInput := parts[0] + "." + parts[1]
	if !hmac.Equal([]byte(parts[2]), []byte(sign(signingInput))) {
		return nil, errors.New("invalid token signature")
	}

	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var claims TokenClaims
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return nil, err
	}

	if time.Now().Unix() > claims.Exp {
		return nil, errors.New("token expired")
	}

	if !claims.Admin {
		return nil, errors.New("insufficient privileges")
	}

	return &claims, nil
}
