package jwt

import (
	"testing"
	"time"
)

func TestJWTSignParse(t *testing.T) {
	mes := Message{
		UserName:  "UserName",
		Email:     "email",
		Phone:     "Phone",
		UserID:    1,
		Timestamp: time.Now().Unix(),
	}

	signed, err := Sign(mes)
	if err != nil {
		t.Fatalf("Failed to sign message: %s", err.Error())
	}

	parsedMessage, err := Parse(signed)
	if err != nil {
		t.Fatalf("Failed to Parse message: %s", err.Error())
	}

	if parsedMessage.UserName != mes.UserName ||
		parsedMessage.Email != mes.Email ||
		parsedMessage.Phone != mes.Phone ||
		parsedMessage.UserID != mes.UserID ||
		parsedMessage.Timestamp != mes.Timestamp {
		t.Fatalf("messages are not equal, mes: %v, parsedMessage: %v", mes, parsedMessage)
	}
}