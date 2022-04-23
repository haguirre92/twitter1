package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	DateBirth time.Time          `bson:"dateBirth" json:"dateBirth,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografia string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SiteWeb   string             `bson:"siteWeb" json:"siteWeb,omitempty"`
}
