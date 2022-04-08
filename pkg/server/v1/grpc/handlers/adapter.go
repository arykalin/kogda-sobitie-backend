package handlers

import (
	"github.com/arykalin/kogda-sobitie-backend/models"
	grpcModels "github.com/arykalin/kogda-sobitie-backend/pkg/server/v1/grpc/models"
)

type adapter struct{}

func (a adapter) authenticateRequest(req *grpcModels.AuthenticateRequest) models.AuthenticateRequest {
	return models.AuthenticateRequest{
		Login:    req.Login,
		Password: req.Password,
	}
}

func (a adapter) authenticateResponse(res models.AuthenticateResponse) *grpcModels.AuthenticateResponse {
	return &grpcModels.AuthenticateResponse{
		Token: res.Token,
	}
}

func (a adapter) createEventRequest(req *grpcModels.CreateEventRequest) models.CreateEventRequest {
	return models.CreateEventRequest{
		Date:        req.Date,
		Title:       req.Title,
		Duration:    req.Duration,
		Link:        req.Link,
		Org:         req.Org,
		Target:      req.Target,
		Where:       req.Where,
		Description: req.Description,
		Amount:      req.Amount,
		Place:       req.Place,
	}
}

func (a adapter) createEventResponse(res models.CreateEventResponse) *grpcModels.CreateEventResponse {
	return &grpcModels.CreateEventResponse{
		Event: &grpcModels.Event{
			Date:        res.Event.Date,
			Title:       res.Event.Title,
			Duration:    res.Event.Duration,
			Link:        res.Event.Link,
			Org:         res.Event.Org,
			Target:      res.Event.Target,
			Where:       res.Event.Where,
			Description: res.Event.Description,
			Amount:      res.Event.Amount,
			Place:       res.Event.Place,
		},
	}
}
