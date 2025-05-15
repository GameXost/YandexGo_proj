package server

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}
