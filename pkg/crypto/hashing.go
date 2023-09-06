package crypto_utils

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (hash []byte, salt []byte, err error) {
	salt = make([]byte, 32)
	_, err = rand.Read(salt)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	hash = argon2.IDKey([]byte(password), salt, 1, 32*1024, 4, 64)
	return hash, salt, nil
}

func VerifyPassword(password string, salt []byte, inHash []byte) (valid bool) {
	hash := argon2.IDKey([]byte(password), salt, 1, 32*1024, 4, 64)
	if bytes.Compare(hash, inHash) != 0 {
		return false
	}
	return true
}
