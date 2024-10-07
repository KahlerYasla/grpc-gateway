package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateJWT(userID string) (string, error) {
	secretKey := []byte("your-secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(secretKey)
}
