package handlers

import (
	"context"

	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
)

type handler struct {
	event
	grpcModels.UnimplementedApiServiceServer
}

//type ApiServiceServer interface {
//	Authenticate(context.Context, *grpcModels.AuthenticateRequest) (*grpcModels.AuthenticateResponse, error)
//	CreateEvent(context.Context, *grpcModels.CreateEventRequest) (*grpcModels.CreateEventResponse, error)
//	GetEvent(context.Context, *grpcModels.GetEventRequest) (*grpcModels.GetEventResponse, error)
//	DeleteEvent(context.Context, *grpcModels.DeleteEventRequest) (*grpcModels.DeleteEventResponse, error)
//	UpdateEvent(context.Context, *grpcModels.UpdateEventRequest) (*grpcModels.UpdateEventResponse, error)
//	ListEvents(context.Context, *grpcModels.ListEventsRequest) (*grpcModels.ListEventsResponse, error)
//}

func (h handler) Authenticate(context.Context, *grpcModels.AuthenticateRequest) (*grpcModels.AuthenticateResponse, error) {
	return nil, nil
}
