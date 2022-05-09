package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTweets struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID      string             `bson:"userid" json:"userId,omitempty"`
	Message     string             `bson:"message" json:"message,omitempty"`
	DateCreated time.Time          `bson:"dateCreated" json:"dateCreated,omitempty"`
}
