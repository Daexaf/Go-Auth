// controllers/transaction_controller.go

package controllers

import (
	"PaymentAPI/initializers"
	"PaymentAPI/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transactionData models.Transaction
	if err := c.ShouldBindJSON(&transactionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var sender, receiver models.User
	if err := initializers.DB.First(&sender, "id = ?", transactionData.SenderAccountID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender"})
		return
	}

	if sender.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender"})
		return
	}

	if err := initializers.DB.First(&receiver, "id = ?", transactionData.ReceiveAccountID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver"})
		return
	}

	if sender.ID == receiver.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sender and receiver cannot be the same"})
		return
	}

	newTransaction := models.Transaction{
		SenderAccountID:  strconv.Itoa(int(sender.ID)),
		ReceiveAccountID: strconv.Itoa(int(receiver.ID)),
		Amount:           transactionData.Amount,
		TimeStamp:        time.Now().Format(time.RFC3339),
	}

	if err := initializers.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction completed"})
}
