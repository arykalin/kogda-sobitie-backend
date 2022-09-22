package clients

import (
	"github.com/arykalin/kogda-sobitie-backend/models"
	generatedClient "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/clients/client"
)

type Client interface {
	Authenticate(request models.AuthenticateRequest) (models.AuthenticateResponse, error)
	CreateEvent(request models.CreateEventRequest) (models.CreateEventResponse, error)
	DeleteEvent(request models.DeleteEventRequest) (models.DeleteEventResponse, error)
	GetEvent(params *APIServiceGetEventParams) (*APIServiceGetEventOK, error)
	ListEvents(params *APIServiceListEventsParams) (*APIServiceListEventsOK, error)
	UpdateEvent(params *APIServiceUpdateEventParams) (*APIServiceUpdateEventOK, error)
}

type client struct {
	grpcCli *generatedClient.Client
}

func NewClient() Client {
	grpcCli := generatedClient.NewHTTPClient()
	return &client{
		grpcCli: grpcCli,
	}
}
