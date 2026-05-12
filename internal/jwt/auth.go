package jwt

import (
	"fmt"
	"go-project/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func InitJWT() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	key := os.Getenv("KEY")
	if key == "" {
		return fmt.Errorf("jwt: KEY environment variable is not set")
	}

	secretKey = []byte(key)
	return nil
}

type Claims struct {
	UserID uint              `json:"user_id"`
	Role   model.AccountType `json:"role"`
	jwt.RegisteredClaims
}

func Generate(user *model.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Role:   user.AccountType,
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
