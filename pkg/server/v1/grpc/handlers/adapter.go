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

//list
func (a adapter) listEventRequest(req *grpcModels.ListEventsRequest) models.ListEventsRequest {
	return models.ListEventsRequest{
		Date: req.Date,
	}
}

func (a adapter) listEventResponse(res models.ListEventsResponse) *grpcModels.ListEventsResponse {
	events := make([]*grpcModels.Event, len(res.Events))
	for i := range res.Events {
		events[i] = &grpcModels.Event{
			Date:        res.Events[i].Date,
			Title:       res.Events[i].Title,
			Duration:    res.Events[i].Duration,
			Link:        res.Events[i].Link,
			Org:         res.Events[i].Org,
			Target:      res.Events[i].Target,
			Where:       res.Events[i].Where,
			Description: res.Events[i].Description,
			Amount:      res.Events[i].Amount,
			Place:       res.Events[i].Place,
		}
	}
	return &grpcModels.ListEventsResponse{
		Events: events,
	}
}

//get
func (a adapter) getEventRequest(req *grpcModels.GetEventRequest) models.GetEventRequest {
	return models.GetEventRequest{
		EventId: req.EventId,
	}
}

func (a adapter) getEventResponse(res models.GetEventResponse) *grpcModels.GetEventResponse {
	return &grpcModels.GetEventResponse{
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

//delete
func (a adapter) deleteEventRequest(req *grpcModels.DeleteEventRequest) models.DeleteEventRequest {
	return models.DeleteEventRequest{
		EventId: req.EventId,
	}
}

func (a adapter) deleteEventResponse(res models.DeleteEventResponse) *grpcModels.DeleteEventResponse {
	return &grpcModels.DeleteEventResponse{
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

//update
func (a adapter) updateEventRequest(req *grpcModels.UpdateEventRequest) models.UpdateEventRequest {
	return models.UpdateEventRequest{
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

func (a adapter) updateEventResponse(res models.UpdateEventResponse) *grpcModels.UpdateEventResponse {
	return &grpcModels.UpdateEventResponse{
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
