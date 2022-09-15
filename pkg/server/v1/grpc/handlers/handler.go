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
	resp, err := h.eventController.GetEvent(ctx, h.adapt.getEventRequest(req))
	if err != nil {
		return nil, err
	}
	return h.adapt.getEventResponse(resp), nil
}

func (h handler) DeleteEvent(ctx context.Context, req *grpcModels.DeleteEventRequest) (*grpcModels.DeleteEventResponse, error) {
	resp, err := h.eventController.DeleteEvent(ctx, h.adapt.deleteEventRequest(req))
	if err != nil {
		return nil, err
	}
	return h.adapt.deleteEventResponse(resp), nil
}

func (h handler) UpdateEvent(ctx context.Context, req *grpcModels.UpdateEventRequest) (*grpcModels.UpdateEventResponse, error) {
	resp, err := h.eventController.UpdateEvent(ctx, h.adapt.updateEventRequest(req))
	if err != nil {
		return nil, err
	}

	return h.adapt.updateEventResponse(resp), nil
}
func (h handler) ListEvents(ctx context.Context, req *grpcModels.ListEventsRequest) (*grpcModels.ListEventsResponse, error) {
	resp, err := h.eventController.ListEvents(ctx, h.adapt.listEventRequest(req))
	if err != nil {
		return nil, err
	}

	return h.adapt.listEventResponse(resp), nil
}

func NewHandler(cntrl eventController.Controller) grpcModels.ApiServiceServer {
	return handler{
		eventController:               cntrl,
		adapt:                         adapter{},
		UnimplementedApiServiceServer: grpcModels.UnimplementedApiServiceServer{},
	}
}
