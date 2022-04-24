package bd

import (
	"context"
	"time"

	"github.com/haguirre92/twitter1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("Users")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		register["lastName"] = user.LastName
	}

	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	register["dateBirth"] = user.DateBirth

	if len(user.Email) > 0 {
		register["email"] = user.Email
	}

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		register["biografia"] = user.Biografia
	}

	if len(user.Ubicacion) > 0 {
		register["ubicacion"] = user.Ubicacion
	}
	if len(user.SiteWeb) > 0 {
		register["siteWeb"] = user.SiteWeb
	}

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
