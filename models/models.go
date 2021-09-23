package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Event Model
type Event struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date          string             `json:"date,omitempty" bson:"date,omitempty" validate:"required,alpha"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty" validate:"required,alpha"`
	Duration      string             `json:"duration,omitempty" bson:"duration,omitempty" validate:"required,alpha"`
	Link          string             `json:"link,omitempty" bson:"link,omitempty" validate:"required,alpha"`
	WhoManages    string             `json:"who_manages,omitempty" bson:"who_manages,omitempty" validate:"required,alpha"`
	ForWhom       string             `json:"for_whom,omitempty" bson:"for_whom,omitempty" validate:"required,alpha"`
	Where         string             `json:"where,omitempty" bson:"where,omitempty" validate:"required,alpha"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty" validate:"required,alpha"`
	WantingPeople string             `json:"wantingPeople,omitempty" bson:"wantingPeople,omitempty" validate:"required,alpha"`
}
