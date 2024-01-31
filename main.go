package main

import (
	"PaymentAPI/controllers"
	"PaymentAPI/initializers"
	"PaymentAPI/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/transaction", middleware.RequireAuth, controllers.CreateTransaction)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)
	r.Run()
}
