package lwr

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
)

type nullTarget struct {
}

func (t *nullTarget) Handle(_ Event) error {
	return nil
}

func GetSign(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	h := sha256.New()
	h.Write(data)
	signed, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signed), nil
}
