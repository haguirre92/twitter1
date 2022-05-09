package bd

import (
	"context"
	"log"
	"time"

	"github.com/haguirre92/twitter1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var allTweets []*models.GetTweets

	condition := bson.M{
		"userid": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "dateCreated", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opts)

	if err != nil {
		log.Fatal(err.Error())
		return allTweets, false
	}

	for cursor.Next(context.TODO()) {
		var register models.GetTweets
		err := cursor.Decode(&register)
		if err != nil {
			return allTweets, false
		}
		allTweets = append(allTweets, &register)
	}
	return allTweets, true
}
