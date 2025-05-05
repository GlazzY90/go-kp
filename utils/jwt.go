package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(uid uint) (string, error) {
  cl := jwt.MapClaims{"user_id": uid, "exp": time.Now().Add(72*time.Hour).Unix()}
  tok := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)
  return tok.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(ts string) (*jwt.Token, error) {
  return jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("JWT_SECRET")), nil
  })
}
