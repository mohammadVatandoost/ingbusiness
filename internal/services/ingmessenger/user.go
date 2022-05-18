package ingmessenger

import (
	"context"
	"github.com/mohammadVatandoost/instabot"
)

func (s *Service) GetUserProfile(token string, userID string) (*instabot.GetUserProfileResponse, error) {
	bot, err := NewClient(token)
	if err != nil {
		return nil, err
	}

	// Getting user profile.
	profile, err := bot.GetUserProfile(
		context.Background(),
		userID,
	)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
