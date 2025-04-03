package utils

import (
	"errors"

	"github.com/google/uuid"
)

func GenerateID() string {
  UUID := uuid.New().String()
  return UUID
}

func ValidateID(id string) error {
  _, err := uuid.Parse(id)
  if err != nil {
    return errors.New("Invalid ID")
  }
  return nil
}
