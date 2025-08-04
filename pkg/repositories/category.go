package repositories

import (
	"context"
	databases "owwi/pkg/database"
	"owwi/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func convertToCategory(categoryRepo models.CategoryRepository) *models.Category {
	return &models.Category{
		ID:          categoryRepo.ID.Hex(),
		User:        categoryRepo.User.Hex(),
		Name:        categoryRepo.Name,
		Description: categoryRepo.Description,
		Type:        categoryRepo.Type.Hex(),
		UsedTime:    categoryRepo.UsedTime,
		CreatedAt:   categoryRepo.CreatedAt,
		UpdatedAt:   categoryRepo.UpdatedAt,
	}
}

func createCategory(category models.CreateCategory) error {
	// Convert the category to a CreateCategoryRepository type
	userId, errP1 := bson.ObjectIDFromHex(category.User)
	if errP1 != nil {
		return errP1
	}

	typeId, errP2 := bson.ObjectIDFromHex(category.Type)
	if errP2 != nil {
		return errP2
	}

	categoryRepo := models.CreateCategoryRepository{
		User:        userId,
		Name:        category.Name,
		Description: category.Description,
		Type:        typeId,
		UsedTime:    category.UsedTime,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	_, err := databases.Client.Collection("categories").InsertOne(context.TODO(), categoryRepo)
	return err
}

func updateCategory(category models.UpdateCategory) error {
	categoryId, errP0 := bson.ObjectIDFromHex(category.ID)
	if errP0 != nil {
		return errP0
	}

	userId, errP1 := bson.ObjectIDFromHex(category.User)
	if errP1 != nil {
		return errP1
	}

	typeId, errP2 := bson.ObjectIDFromHex(category.Type)
	if errP2 != nil {
		return errP2
	}

	categoryRepo := models.CreateCategoryRepository{
		User:        userId,
		Name:        category.Name,
		Description: category.Description,
		Type:        typeId,
		UsedTime:    category.UsedTime,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	_, err := databases.Client.Collection("categories").UpdateByID(
		context.TODO(),
		categoryId,
		map[string]interface{}{
			"$set": categoryRepo,
		},
	)
	return err
}

func getCategoryByID(id string) (*models.Category, error) {
	var categoryRepo models.CategoryRepository
	categoryId, errP := bson.ObjectIDFromHex(id)
	if errP != nil {
		return nil, errP
	}

	err := databases.Client.Collection("categories").FindOne(context.TODO(), bson.M{"_id": categoryId}).Decode(&categoryRepo)
	if err != nil {
		return nil, err
	}

	return convertToCategory(categoryRepo), nil
}

func getAllCategoriesByUser(userID string) ([]models.Category, error) {
	var categories []models.Category
	objectId, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := databases.Client.Collection("categories").Find(context.TODO(), bson.M{"user": objectId})
	if err != nil {
		return nil, err
	}

	var results []models.CategoryRepository
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	for _, categoryRepo := range results {
		categories = append(categories, *convertToCategory(categoryRepo))
	}

	return categories, nil
}

func deleteCategory(id string) error {
	objectId, errP := bson.ObjectIDFromHex(id)
	if errP != nil {
		return errP
	}

	_, err := databases.Client.Collection("categories").DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}

var CategoryRepository = struct {
	CreateCategory         func(models.CreateCategory) error
	UpdateCategory         func(models.UpdateCategory) error
	GetCategoryByID        func(string) (*models.Category, error)
	GetAllCategoriesByUser func(string) ([]models.Category, error)
	DeleteCategory         func(string) error
}{
	CreateCategory:         createCategory,
	UpdateCategory:         updateCategory,
	GetCategoryByID:        getCategoryByID,
	GetAllCategoriesByUser: getAllCategoriesByUser,
	DeleteCategory:         deleteCategory,
}
