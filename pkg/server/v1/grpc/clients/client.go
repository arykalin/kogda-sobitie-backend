package clients

import (
	"fmt"

	"github.com/arykalin/kogda-sobitie-backend/internal/event_controller/models"
	generatedClient "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/client"
	"github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/client/api_service"
	clientModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client interface {
	Authenticate(request models.AuthenticateRequest) (models.AuthenticateResponse, error)
	CreateEvent(request models.CreateEventRequest) (models.CreateEventResponse, error)
	DeleteEvent(request models.DeleteEventRequest) (models.DeleteEventResponse, error)
	GetEvent(request models.GetEventRequest) (models.GetEventResponse, error)
	ListEvents(request models.ListEventsRequest) (models.ListEventsResponse, error)
	UpdateEvent(request models.UpdateEventRequest) (models.UpdateEventResponse, error)
}

type client struct {
	grpcCli *generatedClient.YearWheel
}

func (c client) Authenticate(request models.AuthenticateRequest) (models.AuthenticateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c client) CreateEvent(request models.CreateEventRequest) (models.CreateEventResponse, error) {
	resp, err := c.grpcCli.APIService.APIServiceCreateEvent(&api_service.APIServiceCreateEventParams{
		Context:    nil,
		HTTPClient: nil,
		Body: &clientModels.GrpcCreateEventRequest{
			Amount:      request.Amount,
			Date:        request.Date,
			Description: request.Description,
			Duration:    request.Duration,
			Link:        request.Link,
			Org:         request.Org,
			Place:       request.Place,
			Target:      request.Target,
			Title:       request.Title,
			Where:       request.Where,
		},
	})
	if err != nil {
		return models.CreateEventResponse{}, fmt.Errorf("failed to create event: %w", err)
	}
	return models.CreateEventResponse{
		Event: models.Event{
			ID:          primitive.ObjectID{},
			Date:        resp.Payload.Event.Date,
			Title:       resp.Payload.Event.Title,
			Duration:    resp.Payload.Event.Duration,
			Link:        resp.Payload.Event.Link,
			Org:         resp.Payload.Event.Org,
			Target:      resp.Payload.Event.Title,
			Where:       resp.Payload.Event.Where,
			Description: resp.Payload.Event.Description,
			Amount:      resp.Payload.Event.Amount,
			Place:       resp.Payload.Event.Place,
			Private:     false,
		},
	}, nil
}

func (c client) DeleteEvent(request models.DeleteEventRequest) (models.DeleteEventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c client) GetEvent(request models.GetEventRequest) (models.GetEventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c client) ListEvents(request models.ListEventsRequest) (models.ListEventsResponse, error) {
	resp, err := c.grpcCli.APIService.APIServiceListEvents(&api_service.APIServiceListEventsParams{
		Date: request.Date,
	})
	if err != nil {
		return models.ListEventsResponse{}, fmt.Errorf("failed to list events: %w", err)
	}
	var list []models.Event
	for _, event := range resp.Payload.Events {
		list = append(list, models.Event{
			Date:        event.Date,
			Title:       event.Title,
			Duration:    event.Duration,
			Link:        event.Link,
			Org:         event.Org,
			Target:      event.Target,
			Where:       event.Where,
			Description: event.Description,
			Amount:      event.Amount,
			Place:       event.Place,
			Private:     false,
		})
	}
	return models.ListEventsResponse{
		Events: list,
	}, nil
}

func (c client) UpdateEvent(request models.UpdateEventRequest) (models.UpdateEventResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewClient() Client {
	grpcCli := generatedClient.NewHTTPClient(nil)
	return &client{
		grpcCli: grpcCli,
	}
}
