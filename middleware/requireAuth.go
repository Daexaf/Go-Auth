package middleware

import (
	"PaymentAPI/initializers"
	"PaymentAPI/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		c.Next()
		return
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}

func RequireAuthForTransaction(c *gin.Context) {
	RequireAuth(c)

	user := c.MustGet("user").(models.User)

	var transactionData models.Transaction
	if err := c.ShouldBindJSON(&transactionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		c.Abort()
		return
	}

	// Check if the sender is authenticated
	if strconv.Itoa(int(user.ID)) != transactionData.SenderAccountID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Sender does not match authenticated user"})
		return
	}

	// Check if the receiver is authenticated
	var receiver models.User
	initializers.DB.First(&receiver, transactionData.ReceiveAccountID)
	if receiver.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Receiver not found"})
		return
	}

	c.Next()
}
