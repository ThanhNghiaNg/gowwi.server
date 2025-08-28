package repositories

import (
	"context"
	databases "owwi/pkg/database"
	"owwi/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// func convertToTransaction(transactionRepo models.TransactionRepository) *models.Transaction {
// 	return &models.Transaction{
// 		ID:          transactionRepo.ID.Hex(),
// 		User:        transactionRepo.User.Hex(),
// 		Name:        transactionRepo.Name,
// 		Description: transactionRepo.Description,
// 		Type:        transactionRepo.Type.Hex(),
// 		UsedTime:    transactionRepo.UsedTime,
// 		CreatedAt:   transactionRepo.CreatedAt,
// 		UpdatedAt:   transactionRepo.UpdatedAt,
// 	}
// }

func createTransaction(transaction models.CreateTransaction) error {
	userId, errP1 := bson.ObjectIDFromHex(transaction.User)
	if errP1 != nil {
		return errP1
	}

	typeId, errP2 := bson.ObjectIDFromHex(transaction.TypeID)
	if errP2 != nil {
		return errP2
	}

	categoryId, errP3 := bson.ObjectIDFromHex(transaction.CategoryID)
	if errP3 != nil {
		return errP3
	}

	partnerId, errP3 := bson.ObjectIDFromHex(transaction.PartnerID)
	if errP3 != nil {
		return errP3
	}

	transactionRepo := models.CreateTransactionRepository{
		User:        userId,
		TypeID: typeId,
		TypeName:  transaction.TypeName,
		CategoryID: categoryId,
		CategoryName: transaction.CategoryName,
		PartnerID:   partnerId,
		PartnerName: transaction.PartnerName,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Description: transaction.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	_, err := databases.Client.Collection("transactions").InsertOne(context.TODO(), transactionRepo)
	return err
}

// func updateTransaction(transaction models.UpdateTransaction) error {
// 	transactionId, errP0 := bson.ObjectIDFromHex(transaction.ID)
// 	if errP0 != nil {
// 		return errP0
// 	}

// 	userId, errP1 := bson.ObjectIDFromHex(transaction.User)
// 	if errP1 != nil {
// 		return errP1
// 	}

// 	typeId, errP2 := bson.ObjectIDFromHex(transaction.Type)
// 	if errP2 != nil {
// 		return errP2
// 	}

// 	transactionRepo := models.CreateTransactionRepository{
// 		User:        userId,
// 		Name:        transaction.Name,
// 		Description: transaction.Description,
// 		Type:        typeId,
// 		UsedTime:    transaction.UsedTime,
// 		CreatedAt:   time.Now().UTC(),
// 		UpdatedAt:   time.Now().UTC(),
// 	}

// 	_, err := databases.Client.Collection("transactions").UpdateByID(
// 		context.TODO(),
// 		transactionId,
// 		map[string]interface{}{
// 			"$set": transactionRepo,
// 		},
// 	)
// 	return err
// }

// func getTransactionByID(id string) (*models.Transaction, error) {
// 	var transactionRepo models.TransactionRepository
// 	transactionId, errP := bson.ObjectIDFromHex(id)
// 	if errP != nil {
// 		return nil, errP
// 	}

// 	err := databases.Client.Collection("transactions").FindOne(context.TODO(), bson.M{"_id": transactionId}).Decode(&transactionRepo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return convertToTransaction(transactionRepo), nil
// }

// func getAllTransactionsByUser(userID string) ([]models.Transaction, error) {
// 	var transactions []models.Transaction
// 	objectId, err := bson.ObjectIDFromHex(userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cursor, err := databases.Client.Collection("transactions").Find(context.TODO(), bson.M{"user": objectId})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var results []models.TransactionRepository
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		return nil, err
// 	}

// 	for _, transactionRepo := range results {
// 		transactions = append(transactions, *convertToTransaction(transactionRepo))
// 	}

// 	return transactions, nil
// }

// func deleteTransaction(id string) error {
// 	objectId, errP := bson.ObjectIDFromHex(id)
// 	if errP != nil {
// 		return errP
// 	}

// 	_, err := databases.Client.Collection("transactions").DeleteOne(context.TODO(), bson.M{"_id": objectId})
// 	return err
// }

var TransactionRepository = struct {
	CreateTransaction         func(models.CreateTransaction) error
	UpdateTransaction         func(models.UpdateTransaction) error
	GetTransactionByID        func(string) (*models.Transaction, error)
	DeleteTransaction         func(string) error
}{
	CreateTransaction:         createTransaction,
	// UpdateTransaction:         updateTransaction,
	// GetTransactionByID:        getTransactionByID,
	// DeleteTransaction:         deleteTransaction,
}
