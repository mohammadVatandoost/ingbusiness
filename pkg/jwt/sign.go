package jwt

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
)

func Sign(msg Message) (string, error) {
	data, _ := json.Marshal(msg)
	mapData := make(jwt.MapClaims)
	err := json.Unmarshal(data, &mapData)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapData)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
