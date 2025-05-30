package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const privateKeyPEM = `
-----BEGIN PRIVATE KEY-----
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
61KRSXDzFtGxDRh4UGaxxB/5PKq3zw==
-----END PRIVATE KEY-----


`

func main() {
	// Распарсим приватный ключ
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		panic("failed to decode PEM block")
	}
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	rsaPriv, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		panic("not an RSA private key")
	}

	// Собираем claims
	claims := jwt.MapClaims{
		"sub": "1",                                   // driverID, можешь поменять на любой другой id
		"exp": time.Now().Add(time.Hour * 24).Unix(), // истекает через 24 часа
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(rsaPriv)
	if err != nil {
		panic(err)
	}
	fmt.Println("JWT TOKEN:", signed)
}
