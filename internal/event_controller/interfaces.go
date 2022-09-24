package eventController

import (
	"context"

	"github.com/arykalin/kogda-sobitie-backend/internal/models"
)

type Controller interface {
	Authenticate(ctx context.Context, req models.AuthenticateRequest) (resp models.AuthenticateResponse, err error)
	CreateEvent(ctx context.Context, req models.CreateEventRequest) (resp models.CreateEventResponse, err error)
	ListEvents(ctx context.Context, req models.ListEventsRequest) (resp models.ListEventsResponse, err error)
	GetEvent(ctx context.Context, req models.GetEventRequest) (resp models.GetEventResponse, err error)
	DeleteEvent(ctx context.Context, req models.DeleteEventRequest) (resp models.DeleteEventResponse, err error)
	UpdateEvent(ctx context.Context, req models.UpdateEventRequest) (resp models.UpdateEventResponse, err error)
}
