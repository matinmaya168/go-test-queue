package main

import (
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"payment-queue/db"
	_ "payment-queue/docs"
	"payment-queue/handlers"
	"payment-queue/jobs"
	"payment-queue/middleware"
	"payment-queue/models"
)

// @title Payment Queue API
// @version 1.0
// @description API for managing a payment processing queue with MySQL and Gin
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Payment{})

	r := gin.Default()
	rl := middleware.NewRateLimiter(10, time.Minute) // 10 requests per minute
	r.Use(rl.RateLimiter())

	authorized := r.Group("/")
	{
		authorized.POST("/payments", handlers.EnqueuePayment)
		authorized.GET("/payments", handlers.ListPayments)
		authorized.GET("/payments/:id", handlers.GetPayment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go jobs.ProcessQueue()

	r.Run(":8080")
}
