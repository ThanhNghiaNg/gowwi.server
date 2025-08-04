package api

import (
	"owwi/pkg/models"
	"owwi/pkg/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Create Transaction
// @Description Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param register body models.CreateTransaction true "Transaction data"
// @Security BearerAuth
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /transactions [post]
func createTransaction(c *gin.Context) {
	var transactionData models.CreateTransaction
	if err := c.BindJSON(&transactionData); err != nil || transactionData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := repositories.TransactionRepository.CreateTransaction(transactionData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create type"})
		return
	}

	c.Status(201)
}

// @Summary Update Transaction
// @Description Update Transaction By ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param register body models.UpdateTransaction true "Transaction data"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /transactions [put]
func updateTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Transaction ID is required"})
		return
	}
	var transactionData models.UpdateTransaction
	if err := c.BindJSON(&transactionData); err != nil || transactionData.Name == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	transactionData.ID = id
	if err := repositories.TransactionRepository.UpdateTransaction(transactionData); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.Status(200)
}

// @Summary Get Transaction By ID
// @Description Get Transaction By ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /transactions/:id [get]
func getTransactionByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Transaction ID is required"})
		return
	}

	transaction, err := repositories.TransactionRepository.GetTransactionByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve transaction", "err": err.Error()})
		return
	}

	if transaction == nil {
		c.JSON(404, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(200, transaction)
}

// @Summary Get All Transactions By User
// @Description Get all transactions associated with the authenticated user
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /transactions [get]
func getAllTransactionsByUser(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	transactions, err := repositories.TransactionRepository.GetAllTransactionsByUser(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	if len(transactions) == 0 {
		c.JSON(404, gin.H{"message": "No transactions found for this user", "user_id": userID, "transactions": transactions})
		return
	}

	c.JSON(200, transactions)
}
// @Summary Delete Transaction
// @Description Delete Transaction By ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Security BearerAuth
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /transactions/:id [delete]
func deleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Transaction ID is required"})
		return
	}

	if err := repositories.TransactionRepository.DeleteTransaction(id); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.Status(204)
}

var TransactionApi = struct {
	CreateTransaction         func(*gin.Context)
	UpdateTransaction         func(*gin.Context)
	GetTransactionByID        func(*gin.Context)
	DeleteTransaction         func(*gin.Context)
}{
	CreateTransaction:         createTransaction,
	UpdateTransaction:         updateTransaction,
	GetTransactionByID:        getTransactionByID,
	DeleteTransaction:         deleteTransaction,
}
