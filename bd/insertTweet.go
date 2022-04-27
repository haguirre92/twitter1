package bd

import (
	"context"
	"time"

	"github.com/haguirre92/twitter1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	register := bson.M{
		"userid":      tweet.UserID,
		"message":     tweet.Message,
		"dateCreated": tweet.DateCreated,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return string(""), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
