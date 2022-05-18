package ingmessenger

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/ingaccounts"
	"github.com/mohammadVatandoost/ingbusiness/internal/savedmessages"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

type Service struct {
	usersDirectory         users.Querier
	ingAccountsDirectory   ingaccounts.Querier
	savedMessagesDirectory savedmessages.Querier
}

func New(usersDirectory users.Querier, ingAccountsDirectory ingaccounts.Querier,
	savedMessagesDirectory savedmessages.Querier) *Service {
	return &Service{
		usersDirectory:         usersDirectory,
		ingAccountsDirectory:   ingAccountsDirectory,
		savedMessagesDirectory: savedMessagesDirectory,
	}
}
