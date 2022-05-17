package ingmessenger

import "github.com/mohammadVatandoost/instabot"

func NewClient(token string) (*instabot.Client, error) {
	return instabot.New(token)
}
