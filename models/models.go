package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event Model
type Event struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date          string             `json:"date,omitempty" bson:"date,omitempty" validate:"required"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Duration      string             `json:"duration,omitempty" bson:"duration,omitempty" validate:"required"`
	Link          string             `json:"link,omitempty" bson:"link,omitempty" validate:"required"`
	WhoManages    string             `json:"who_manages,omitempty" bson:"who_manages,omitempty" validate:"required"`
	ForWhom       string             `json:"for_whom,omitempty" bson:"for_whom,omitempty" validate:"required"`
	Where         string             `json:"where,omitempty" bson:"where,omitempty" validate:"required"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty" validate:"required"`
	WantingPeople string             `json:"wanting_people,omitempty" bson:"wanting_people,omitempty" validate:"required"`
}
