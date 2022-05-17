package ingmessenger

import (
	"context"
	"github.com/mohammadVatandoost/instabot"
)

//ToDo: send message with image and sticker

func (m *Messenger) SendTextMessage(token string, receiverID string, message string) error {
	bot, err := NewClient(token)
	if err != nil {
		return err
	}

	_, err = bot.SendMessage(
		context.Background(),
		receiverID,
		instabot.NewTextMessage(message),
	)
	if err != nil {
		return err
	}

	return nil
}
