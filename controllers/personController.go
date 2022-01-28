package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/arykalin/kogda-sobitie-backend/auth"
	"github.com/arykalin/kogda-sobitie-backend/db"
	middlewares "github.com/arykalin/kogda-sobitie-backend/handlers"
	"github.com/arykalin/kogda-sobitie-backend/models"
	"github.com/arykalin/kogda-sobitie-backend/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Dbconnect()

// Auths -> get token
var Auths = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//TODO: generate token only for valid google users
	//https://qvault.io/golang/how-to-implement-sign-in-with-google-in-golang/
	idToken := request.URL.Query().Get("id_token")

	claims, err := auth.ValidateGoogleJWT(idToken)
	if err != nil {
		middlewares.ErrorResponse(fmt.Sprintf("Invalid token"), response)
		return
	}
	if !claims.EmailVerified {
		middlewares.ErrorResponse("Email is not verified", response)
		return
	}

	validToken, err := middlewares.GenerateJWT(claims.Email)
	if err != nil {
		middlewares.ErrorResponse("Failed to generate token", response)
		return
	}

	info := models.UserInfo{
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		Name:          fmt.Sprintf("%s %s", claims.FirstName, claims.LastName),
		Picture:       "",
		GivenName:     claims.FirstName,
		FamilyName:    claims.LastName,
	}
	middlewares.UserInfoResponse(info, validToken, response)
})

// CreateEventEndpoint -> create event
var CreateEventEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var event models.Event
	err := json.NewDecoder(request.Body).Decode(&event)
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
})

// GetEventsEndpoint -> get events
var GetEventsEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
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
})

// GetEventEndpoint -> get event by id
var GetEventEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var event models.Event

	collection := client.Database("golang").Collection("events")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&event)
	if err != nil {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	middlewares.SuccessRespond(event, response)
})

// DeleteEventEndpoint -> delete event by id
var DeleteEventEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
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
})

// UpdateEventEndpoint -> update event by id
var UpdateEventEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var (
		event    models.Event
		oldEvent models.Event
	)

	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(event); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}

	collection := client.Database("golang").Collection("events")
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&oldEvent)
	if err != nil {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	data, err := bson.Marshal(event)
	if err != nil {
		middlewares.ErrorResponse("Failed to marshal event", response)
		return
	}
	res, err := collection.ReplaceOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: id}},
		data,
	)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if res.MatchedCount == 0 {
		middlewares.ErrorResponse("Event does not exist", response)
		return
	}
	middlewares.SuccessResponse("Updated", response)
})

// UploadFileEndpoint -> upload file
var UploadFileEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("file")
	// fileName := request.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile("uploaded/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	middlewares.SuccessResponse("Uploaded Successfully", response)
})
