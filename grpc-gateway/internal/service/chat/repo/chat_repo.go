package repo

import (
	"context"
	"time"

	"src/internal/service/chat/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var message_rooms_collection = client.Database("dbname").Collection("message_rooms")

func InitMongoClient(uri string) error {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func AddMessage(roomID string, message model.Message) error {
	filter := bson.M{"_id": roomID}
	update := bson.M{"$push": bson.M{"messages:": message}}

	_, err := message_rooms_collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func GetMessages(roomID string) ([]*model.Message, error) {
	filter := bson.M{"_id": roomID}

	var result struct {
		Messages []*model.Message `bson:"messages"`
	}

	err := message_rooms_collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Messages, nil
}
