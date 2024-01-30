package controllers

import (
	"PaymentAPI/initializers"
	"PaymentAPI/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	//get data from req body

	//create transaction
	var transactionData models.Transaction
	if err := c.ShouldBindJSON(&transactionData); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	transaction := models.Transaction{
		SenderAccountID:  transactionData.SenderAccountID,
		ReceiveAccountID: transactionData.ReceiveAccountID,
		Amount:           transactionData.Amount,
		TimeStamp:        transactionData.TimeStamp,
	}

	result := initializers.DB.Create(&transaction)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	//return it
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
