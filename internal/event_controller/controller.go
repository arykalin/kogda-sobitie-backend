package eventController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/arykalin/kogda-sobitie-backend/internal/auth"
	"github.com/arykalin/kogda-sobitie-backend/models"
	"github.com/arykalin/kogda-sobitie-backend/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type controller struct {
	logger     *zap.SugaredLogger
	dbClient   *mongo.Client
	collection *mongo.Collection
}

func (c *controller) Authenticate(req models.AuthenticateRequest) (resp models.AuthenticateResponse, err error) {
	//TODO: generate token only for valid google users
	//https://qvault.io/golang/how-to-implement-sign-in-with-google-in-golang/

	claims, err := auth.ValidateGoogleJWT(req.GoogleToken)
	if err != nil {
		return
	}
	if !claims.EmailVerified {
		return resp, fmt.Errorf("email is not verified")
	}

	validToken, err := middlewares.GenerateJWT(claims.Email)
	if err != nil {
		return resp, fmt.Errorf("failed to generate token: %w", err)
	}

	info := models.Account{
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		Name:          fmt.Sprintf("%s %s", claims.FirstName, claims.LastName),
		Picture:       "",
		GivenName:     claims.FirstName,
		FamilyName:    claims.LastName,
	}
	resp.Account = info
	resp.Token = validToken
	return resp, nil
}

// CreateEvent -> create event
func (c *controller) CreateEvent(req models.CreateEventRequest) (resp models.CreateEventResponse, err error) {
	event := models.Event{
		ID:          primitive.ObjectID{},
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

	if ok, errors := validators.ValidateInputs(event); !ok {
		return resp, fmt.Errorf("validation error: %s", errors)
	}

	result, err := c.collection.InsertOne(context.TODO(), event)
	if err != nil {
		return resp, fmt.Errorf("failed to insert event: %w", err)
	}
	res, _ := json.Marshal(result.InsertedID)
	c.logger.Debugf("Inserted event: %s", res)
	return resp, nil
}

// ListEvents -> get events
func (c *controller) ListEvents(req models.ListEventsRequest) (resp models.ListEventsResponse, err error) {
	var events []*models.Event

	cursor, err := c.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return resp, fmt.Errorf("failed to find events: %w", err)
	}
	for cursor.Next(context.TODO()) {
		var event models.Event
		err := cursor.Decode(&event)
		if err != nil {
			log.Fatal(err)
		}

		events = append(events, &event)
	}
	if err := cursor.Err(); err != nil {
		return resp, fmt.Errorf("failed to decode events: %w", err)
	}
	resp.Events = events
	return
}

// GetEvent -> get event by id
func (c *controller) GetEvent(req models.GetEventRequest) (resp models.GetEventResponse, err error) {
	eventID, err := primitive.ObjectIDFromHex(req.EventId)
	err = c.collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: eventID}}).Decode(&resp.Event)
	if err != nil {
		return resp, fmt.Errorf("failed to find event: %w", err)
	}
	return resp, nil
}

// DeleteEventEndpoint -> delete event by id
func (c *controller) DeleteEventEndpoint(req models.DeleteEventRequest) (resp models.DeleteEventResponse, err error) {
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var event models.Event

	collection := client.Database("golang").Collection("events")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&event)
	if err != nil {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	_, derr := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if derr != nil {
		middlewares.ServerErrResponse(derr.Error(), response)
		return
	}
	middlewares.SuccessResponse("Deleted", response)
}

// UpdateEventEndpoint -> update event by id
func (c *controller) UpdateEventEndpoint(req models.UpdateEventRequest) (resp models.UpdateEventResponse, err error) {
	id, _ := primitive.ObjectIDFromHex(req.EventId)
	var (
		event = models.Event{
			ID:          id,
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
		oldEvent models.Event
	)

	if ok, errors := validators.ValidateInputs(event); !ok {
		return resp, fmt.Errorf("validation errors: %v", errors)
	}

	collection := client.Database("golang").Collection("events")
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&oldEvent)
	if err != nil {
		return resp, fmt.Errorf("event does not exist: %w", err)
	}
	data, err := bson.Marshal(event)
	if err != nil {
		return resp, fmt.Errorf("failed to marshal event: %w", err)
	}
	res, err := collection.ReplaceOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: id}},
		data,
	)
	if err != nil {
		return resp, fmt.Errorf("failed to update event: %w", err)
	}
	if res.MatchedCount == 0 {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	middlewares.SuccessResponse("Updated", response)
	return
}

func NewController(
	dbClient *mongo.Client,
	collection *mongo.Collection,
	logger *zap.SugaredLogger,
) *controller {
	return &controller{
		dbClient:   dbClient,
		collection: collection,
		logger:     logger,
	}
}
