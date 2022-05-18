package frequentmessages

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/savedmessages"
)

type Service struct {
	savedMessagesDirectory savedmessages.Querier
}

func New(savedMessagesDirectory savedmessages.Querier) *Service {
	return &Service{
		savedMessagesDirectory: savedMessagesDirectory,
	}
}
