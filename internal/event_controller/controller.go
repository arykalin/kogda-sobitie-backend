package eventController

import (
	"context"
	"encoding/json"
	"fmt"

	"log"

	"strings"

	"github.com/arykalin/kogda-sobitie-backend/db"
	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/arykalin/kogda-sobitie-backend/internal/auth"
	"github.com/arykalin/kogda-sobitie-backend/models"
	"github.com/arykalin/kogda-sobitie-backend/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client = db.Dbconnect()

type controller struct {
	dbClient *mongo.Client
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
	middlewares.UserInfoResponse(info, validToken, response)
	return
}

// CreateEventEndpoint -> create event
func (c *controller) CreateEventEndpoint(req models.CreateEventRequest) (resp models.CreateEventResponse, err error) {
	var event models.Event
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(event); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}
	collection := client.Database("golang").Collection("events")
	result, err := collection.InsertOne(context.TODO(), event)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
}

// GetEventsEndpoint -> get events
func (c *controller) GetEventsEndpoint(req models.ListEventsRequest) (resp models.ListEventsResponse, err error) {
	var events []*models.Event

	collection := client.Database("golang").Collection("events")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
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
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessArrRespond(events, response)
}

// GetEventEndpoint -> get event by id
func (c *controller) GetEventEndpoint(req models.GetEventRequest) (resp models.GetEventResponse, err error) {
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var event models.Event

	collection := client.Database("golang").Collection("events")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&event)
	if err != nil {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	middlewares.SuccessRespond(event, response)
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

func NewController(dbClient *mongo.Client) *controller {
	return &controller{dbClient: dbClient}
}
