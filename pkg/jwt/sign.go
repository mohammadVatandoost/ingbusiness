package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func Sign(msg Message) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        strconv.Itoa(msg.UserID),
		"username":  msg.UserName,
		"phone":     msg.Phone,
		"email":     msg.Email,
		"timestamp": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
