package repositories

import (
	"context"
	"owwi/pkg/database"
	"owwi/pkg/models"
)



func createType(_type models.Type) error {
	_, err := databases.Client.Collection("types").InsertOne(context.TODO(), _type)
	return err
}

var TypeRepository = struct{
	CreateType func(models.Type) error
}{
	CreateType: createType,
}
