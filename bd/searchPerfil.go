package bd

import (
	"context"
	"log"
	"time"

	"github.com/haguirre92/twitter1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchPerfil(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("Users")

	var perfil models.User
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectID,
	}

	err := col.FindOne(ctx, condition).Decode(&perfil)

	perfil.Password = ""
	if err != nil {
		log.Fatalln("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
