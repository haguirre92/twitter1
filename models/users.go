package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omiteempty" json:"id"`
	Name      string             `bson:"name" json:"name,omiteempty"`
	LastName  string             `bson:"lastName" json:"lastName,omiteempty"`
	DateBirth time.Time          `bson:"dateBirth" json:"dateBirth,omiteempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omiteempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omiteempty"`
	Banner    string             `bson:"banner" json:"banner,omiteempty"`
	Biografia string             `bson:"biografia" json:"biografia,omiteempty"`
	Ubicacion string             `bson:"ubicacion" json:"ubicacion,omiteempty"`
	SiteWeb   string             `bson:"siteWeb" json:"siteWeb,omiteempty"`
}
