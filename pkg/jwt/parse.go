package jwt

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
)

var hmacSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

func Parse(signedMessage string) (*Message, error) {
	token, err := jwt.Parse(signedMessage, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		j, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}

		msg := &Message{}

		err = json.Unmarshal(j, msg)
		if err != nil {
			return nil, err
		}

		return msg, nil
	}

	return nil, fmt.Errorf("token is not valid")
}
