package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event Model
type Event struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date        string             `json:"date,omitempty" bson:"date,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Duration    string             `json:"duration,omitempty" bson:"duration,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Org         string             `json:"org,omitempty" bson:"org,omitempty"`
	Target      string             `json:"target,omitempty" bson:"target,omitempty"`
	Where       string             `json:"where,omitempty" bson:"where,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Amount      string             `json:"amount,omitempty" bson:"amount,omitempty"`
	Place       string             `json:"place,omitempty" bson:"place,omitempty"`
}

type Account struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email         string             `json:"email" bson:"email"`
	EmailVerified bool               `json:"emailVerified" bson:"emailVerified"`
	Name          string             `json:"name" bson:"name"`
	Picture       string             `json:"picture" bson:"picture"`
	GivenName     string             `json:"givenName" bson:"givenName"`
	FamilyName    string             `json:"familyName" bson:"familyName"`
	Role          string             `json:"role" bson:"role"`
}

type Entity string

const AccountEntity = "ACCOUNT"
const EventEntity = "EVENT"
