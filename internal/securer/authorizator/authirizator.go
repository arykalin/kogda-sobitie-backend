package authorizator

import (
	"context"

	"github.com/arykalin/kogda-sobitie-backend/internal/event_controller/models"
)

type Authorizator interface {
	Authorize(ctx context.Context, user models.Account, entity models.Entity, verb models.RoleVerb) error
}

func NewAuthorizator() Authorizator {
	return nil
}
