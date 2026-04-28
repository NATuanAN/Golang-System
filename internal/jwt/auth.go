package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func init() {
	key := os.Getenv("KEY")
	if key == "" {
		panic("jwt: KEY environment variable is not set")
	}
	secretKey = []byte(key)
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func Generate(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("jwt.Generate: %w", err)
	}

	return signed, nil
}

func Parse(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt.Parse: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("jwt.Parse: invalid token")
	}

	return claims, nil
}
