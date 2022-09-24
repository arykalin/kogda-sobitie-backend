package authorizator

import (
	"context"

	models2 "github.com/arykalin/kogda-sobitie-backend/internal/models"
)

type Authorizator interface {
	Authorize(ctx context.Context, user models2.Account, entity models2.Entity, verb models2.RoleVerb) error
}

func NewAuthorizator() Authorizator {
	return nil
}
