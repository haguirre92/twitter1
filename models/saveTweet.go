package models

import "time"

type SaveTweet struct {
	UserID      string    `bson:"userid" json:"userid, omitempty"`
	Message     string    `bson:"message" json:"message, omitempty"`
	DateCreated time.Time `bson:"dateCreated" json:"dateCreated, omitempty"`
}
