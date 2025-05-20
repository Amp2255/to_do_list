package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tasks struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Status      string             `json:"status,omitempty" validate:"required"`
	Priority    string             `json:"priority,omitempty" validate:"required"`
	//DueDate     primitive.DateTime `json:"duedate,omitempty"`
	DueDate   time.Time `bson:"due_date,omitempty" json:"due_date,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
