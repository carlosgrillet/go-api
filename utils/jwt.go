package utils

import (
	"time"
  "errors"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "3>=bMVe_a_0]40BMX=.m7!z|;CT1PPf"

func GenerateToken(id, email string) (string, error) {
  claims := jwt.MapClaims{
    "userId": id,
    "email": email,
    "expire": time.Now().Add(2 * time.Hour).Unix(),
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (string, error) {
  parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(any, error) {
    _, ok := token.Method.(*jwt.SigningMethodHMAC)
    if !ok {
      return nil, errors.New("Unexpected signing method")
    }
    return []byte(secretKey), nil
  })

  if err != nil {
    return "", err
  }

  if !parsedToken.Valid {
    return "", errors.New("Invalid token.")
  }

  claims, ok := parsedToken.Claims.(jwt.MapClaims)
  if !ok {
    return "", errors.New("Invalid otken claims.")
  }
  
  userId := claims["userId"].(string)

  return userId, nil
}
