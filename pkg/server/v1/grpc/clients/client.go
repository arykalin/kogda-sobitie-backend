package clients

import (
	"github.com/arykalin/kogda-sobitie-backend/models"
	generatedClient "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/client"
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
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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
