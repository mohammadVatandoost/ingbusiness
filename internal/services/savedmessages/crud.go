package savedmessages

import (
	"context"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/savedmessages/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/savedmessages"
)

func (s *Store) Add(ctx context.Context, savedMessage *v1.SavedMessage,
	userID int32) (savedmessages.SavedMessage, error) {
	SavedMessage, err := s.savedMessagesDirectory.AddSavedMessage(ctx,
		savedmessages.AddSavedMessageParams{
			Message:      savedMessage.MessageContent,
			IngAccountID: savedMessage.IngAccountID,
			WriterID:     userID,
		})
	return SavedMessage, err
}

func (s *Store) Delete(ctx context.Context, savedMessage *v1.SavedMessage) error {
	_, err := s.savedMessagesDirectory.DeleteSavedMessage(ctx, savedMessage.Id)
	return err
}

func (s *Store) UpdateSavedMessageMessage(ctx context.Context, savedMessage *v1.SavedMessage) error {
	_, err := s.savedMessagesDirectory.UpdateSavedMessageMessage(ctx,
		savedmessages.UpdateSavedMessageMessageParams{
			ID:      savedMessage.Id,
			Message: savedMessage.MessageContent,
		})
	return err
}

//ToDo: add pagination
func (s *Store) GetSavedMessages(ctx context.Context) ([]savedmessages.SavedMessage, error) {
	SavedMessages, err := s.savedMessagesDirectory.GetSavedMessages(ctx)
	return SavedMessages, err
}
