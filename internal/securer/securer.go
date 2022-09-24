package securer

import (
	"context"
	"fmt"

	"github.com/arykalin/kogda-sobitie-backend/internal/event_controller/models"
	"github.com/arykalin/kogda-sobitie-backend/internal/securer/authenticator"
	"github.com/arykalin/kogda-sobitie-backend/internal/securer/authorizator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type securer struct {
	authenticator.Authenticator
	authorizator.Authorizator
}

type Securer interface {
	Secure(ctx context.Context, name models.Entity, create models.RoleVerb) error
}

func (s *securer) Secure(ctx context.Context, entity models.Entity, verb models.RoleVerb) error {
	user, token, err := s.getToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get token from context: %w", err)
	}
	acc, err := s.Authenticate(ctx, user, token)
	if err != nil {
		return fmt.Errorf("failed to authenticat token: %w", err)
	}

	isAdmin, err := s.accIsAdmin(ctx, acc)
	if err != nil {
		return fmt.Errorf("failed to check admin status: %w", err)
	}
	if isAdmin {
		// skip authorizator for admin user
		return nil
	}

	err = s.Authorize(ctx, acc, entity, verb)
	if err != nil {
		return fmt.Errorf("failed to authorize account %s err: %w", acc.Name, err)
	}

	return nil
}

func (s *securer) getToken(ctx context.Context) (user, token string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", status.Errorf(codes.Unauthenticated, "failed to get metadata from context")
	}

	values := md[AuthUserKey]
	if len(values) == 0 {
		return "", "", status.Errorf(codes.Unauthenticated, "user key is not provided")
	}
	user = values[0]

	values = md[AuthTokenKey]
	if len(values) == 0 {
		return "", "", status.Errorf(codes.Unauthenticated, "token key is not provided")
	}
	token = values[0]
	return user, token, nil
}

func (s *securer) accIsAdmin(ctx context.Context, acc models.Account) (bool, error) {
	//TODO: check if acc have admin role @arykalin CL-259
	return false, nil
}

func NewSecurer() Securer {
	return &securer{
		Authenticator: authenticator.NewAuthenticator(),
		Authorizator:  authorizator.NewAuthorizator(),
	}
}

func MakeSecureContext(ctx context.Context, user, token string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, AuthTokenKey, token)
	ctx = metadata.AppendToOutgoingContext(ctx, AuthUserKey, user)
	return ctx
}
