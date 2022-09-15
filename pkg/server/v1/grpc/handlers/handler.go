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

// TODO: implement methods from controller
func (h handler) CreateEvent(ctx context.Context, req *grpcModels.CreateEventRequest) (*grpcModels.CreateEventResponse, error) {
	resp, err := h.eventController.CreateEvent(ctx, h.adapt.createEventRequest(req))
	if err != nil {
		return nil, err
	}
	return h.adapt.createEventResponse(resp), nil
}

func (h handler) GetEvent(ctx context.Context, req *grpcModels.GetEventRequest) (*grpcModels.GetEventResponse, error) {
	return nil, nil
}
func (h handler) DeleteEvent(ctx context.Context, req *grpcModels.DeleteEventRequest) (*grpcModels.DeleteEventResponse, error) {
	return nil, nil
}
func (h handler) UpdateEvent(ctx context.Context, req *grpcModels.UpdateEventRequest) (*grpcModels.UpdateEventResponse, error) {
	return nil, nil
}
func (h handler) ListEvents(ctx context.Context, req *grpcModels.ListEventsRequest) (*grpcModels.ListEventsResponse, error) {
	return nil, nil
}

func NewHandler() grpcModels.ApiServiceServer {
	return handler{}
}
