package authenticator

import (
	"context"

	"github.com/arykalin/kogda-sobitie-backend/models"
)

type Authenticator interface {
	Authenticate(ctx context.Context, user, token string) (accountModel models.Account, err error)
}
