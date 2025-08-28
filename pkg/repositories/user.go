package repositories

import (
	"context"
	databases "owwi/pkg/database"
	"owwi/pkg/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func convertToUser(user models.UserRepository) *models.User {
	return &models.User{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		FullName:  user.FullName,
		Phone:     user.Phone,
		Address:   user.Address,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func createUser(user models.CreateUser) (bson.ObjectID, error) {
	res, err := databases.Client.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return bson.NilObjectID, err
	}
	// fmt.Println("Inserted ID:", res.InsertedID)
	id := res.InsertedID.(bson.ObjectID)
	return id, err
}

func getUserByID(id string) (*models.User, error) {
	var user models.UserRepository
	userId, errP := bson.ObjectIDFromHex(id)
	if errP != nil {
		return nil, errP
	}
	err := databases.Client.Collection("users").FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return convertToUser(user), nil
}

func getUserByUsername(username string) (*models.User, error) {
	var user models.UserRepository
	err := databases.Client.Collection("users").FindOne(context.TODO(), map[string]string{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return convertToUser(user), nil
}

func updateUser(user models.UpdateUser) error {
	_, err := databases.Client.Collection("users").UpdateByID(
		context.TODO(),
		user.ID,
		map[string]interface{}{
			"$set": user,
		},
	)
	return err
}

func deleteUser(id string) error {
	_, err := databases.Client.Collection("users").DeleteOne(context.TODO(), map[string]string{"id": id})
	return err
}

var UserRepository = struct {
	CreateUser        func(user models.CreateUser) (bson.ObjectID, error)
	GetUserByID       func(id string) (*models.User, error)
	GetUserByUsername func(username string) (*models.User, error)
	UpdateUser        func(user models.UpdateUser) error
	DeleteUser        func(id string) error
}{
	CreateUser:        createUser,
	GetUserByID:       getUserByID,
	GetUserByUsername: getUserByUsername,
	UpdateUser:        updateUser,
	DeleteUser:        deleteUser,
}
