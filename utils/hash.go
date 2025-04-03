package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(text string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(text), 16)
  return string(bytes), err
}

func CompareWithHash(text, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
  if err != nil {
    return false
  }
  return true
}
