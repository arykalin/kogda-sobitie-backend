package authenticator

import (
	"context"

	"github.com/arykalin/kogda-sobitie-backend/internal/event_controller/models"
)

type Authenticator interface {
	Authenticate(ctx context.Context, user, token string) (accountModel models.Account, err error)
}

func NewAuthenticator() Authenticator {
	return nil
}
