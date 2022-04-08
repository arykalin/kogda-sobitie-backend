package handlers

import (
	"context"

	eventController "github.com/arykalin/kogda-sobitie-backend/internal/event_controller"
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
)

type handler struct {
	eventController eventController.Controller
	adapt           adapter
	grpcModels.UnimplementedApiServiceServer
}

func (h handler) Authenticate(ctx context.Context, req *grpcModels.AuthenticateRequest) (*grpcModels.AuthenticateResponse, error) {
	resp, err := h.eventController.Authenticate(ctx, h.adapt.authenticateRequest(req))
	return h.adapt.authenticateResponse(resp), err
}
func (h handler) CreateEvent(context.Context, *grpcModels.CreateEventRequest) (*grpcModels.CreateEventResponse, error) {
	return nil, nil
}
func (h handler) GetEvent(context.Context, *grpcModels.GetEventRequest) (*grpcModels.GetEventResponse, error) {
	return nil, nil
}
func (h handler) DeleteEvent(context.Context, *grpcModels.DeleteEventRequest) (*grpcModels.DeleteEventResponse, error) {
	return nil, nil
}
func (h handler) UpdateEvent(context.Context, *grpcModels.UpdateEventRequest) (*grpcModels.UpdateEventResponse, error) {
	return nil, nil
}
func (h handler) ListEvents(context.Context, *grpcModels.ListEventsRequest) (*grpcModels.ListEventsResponse, error) {
	return nil, nil
}

func NewHandler() grpcModels.ApiServiceServer {
	return handler{}
}
