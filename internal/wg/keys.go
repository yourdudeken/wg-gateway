package wg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

type KeyPair struct {
	Private string
	Public  string
}

func GenerateKeyPair() (KeyPair, error) {
	var privateKey [32]byte
	_, err := rand.Read(privateKey[:])
	if err != nil {
		return KeyPair{}, fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Clamp the private key as per Curve25519 spec
	privateKey[0] &= 248
	privateKey[31] &= 127
	privateKey[31] |= 64

	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)

	return KeyPair{
		Private: base64.StdEncoding.EncodeToString(privateKey[:]),
		Public:  base64.StdEncoding.EncodeToString(publicKey[:]),
	}, nil
}
