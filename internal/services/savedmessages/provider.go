package savedmessages

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/savedmessages"
)

type Store struct {
	savedMessagesDirectory savedmessages.Querier
}

func New(savedMessagesDirectory savedmessages.Querier) *Store {
	return &Store{
		savedMessagesDirectory: savedMessagesDirectory,
	}
}
