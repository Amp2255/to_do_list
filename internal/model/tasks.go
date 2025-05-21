package model

import (
	"to_do_list/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tasks struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	Status      string             `json:"status,omitempty" validate:"required"`
	Priority    string             `json:"priority,omitempty" validate:"required"`
	//DueDate     primitive.DateTime `json:"duedate,omitempty"`
	DueDate utils.DateOnly `bson:"due_date" json:"due_date"`
}
