package models

import "time"

type SaveTweet struct {
	UserID      string    `json:"userid" json:"userid, omitempty"`
	Message     string    `json:"message" json:"message, omitempty"`
	DateCreated time.Time `json:"dateCreated" json:"dateCreated, omitempty"`
}
