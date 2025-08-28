package repositories

import (
	"context"
	databases "owwi/pkg/database"
	"owwi/pkg/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func createType(_type models.Type) error {
	_, err := databases.Client.Collection("types").InsertOne(context.TODO(), _type)
	return err
}

func getTypeByID(id string) (*models.Type, error) {
	var typeRepo models.Type
	typeId, _ := bson.ObjectIDFromHex(id)
	err := databases.Client.Collection("types").FindOne(context.TODO(), bson.M{"_id": typeId}).Decode(&typeRepo)
	if err != nil {
		return nil, err
	}
	return &typeRepo, nil
}

func getTypesByUserID(id string) ([]*models.Type, error) {
	var types []*models.Type
	cursor, err := databases.Client.Collection("types").Find(context.TODO(), bson.M{"user": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var typeRepo models.Type
		if err := cursor.Decode(&typeRepo); err != nil {
			return nil, err
		}
		types = append(types, &typeRepo)
	}
	return types, nil
}

var TypeRepository = struct {
	CreateType  func(models.Type) error
	GetTypeByID func(string) (*models.Type, error)
	GetTypesByUserID func(string) ([]*models.Type, error)
}{
	CreateType:  createType,
	GetTypeByID: getTypeByID,
	GetTypesByUserID: getTypesByUserID,
}
