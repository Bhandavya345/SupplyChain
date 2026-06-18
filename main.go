// @title Supply Chain Management API
// @version 1.0
// @description Supply Chain Management System using Golang, Gin, PostgreSQL
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	_ "github.com/Bhandavya345/supply-chain-system/docs"
	"github.com/Bhandavya345/supply-chain-system/internal/auth"
	"github.com/Bhandavya345/supply-chain-system/internal/database"
	"github.com/Bhandavya345/supply-chain-system/internal/inventory"
	"github.com/Bhandavya345/supply-chain-system/internal/middleware"
	"github.com/Bhandavya345/supply-chain-system/internal/shipment"
	"github.com/Bhandavya345/supply-chain-system/internal/supplier"
	"github.com/Bhandavya345/supply-chain-system/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Connect Database
	database.ConnectDB()

	// Auto Migrate Tables
	database.DB.AutoMigrate(
		&models.User{},
		&models.Inventory{},
		&models.Supplier{},
		&models.Shipment{},
	)

	// Create Gin Router

	router := gin.Default()

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Supply Chain Management API Running",
		})
	})

	// =========================
	// Authentication Routes
	// =========================

	router.POST("/api/auth/register", auth.Register)
	router.POST("/api/auth/login", auth.Login)

	// =========================
	// Protected Routes
	// =========================

	api := router.Group("/api")
	//api.Use(middleware.AuthMiddleware())

	// Inventory Routes
	api.POST("/inventory", inventory.Create)
	api.GET("/inventory", inventory.GetAll)
	api.GET("/inventory/:id", inventory.GetByID)
	api.PUT("/inventory/:id", inventory.Update)
	api.DELETE("/inventory/:id", inventory.Delete)

	// Supplier Routes
	api.POST("/suppliers", supplier.Create)
	api.GET("/suppliers", supplier.GetAll)
	api.GET("/suppliers/:id", supplier.GetByID)
	api.PUT("/suppliers/:id", supplier.Update)
	api.DELETE("/suppliers/:id", supplier.Delete)
	// Shipment Routes
	api.POST("/shipments", shipment.Create)
	api.GET("/shipments", shipment.GetAll)
	api.GET("/shipments/:id", shipment.GetByID)
	api.PUT("/shipments/:id", shipment.Update)
	api.DELETE("/shipments/:id", shipment.Delete)
	api.GET("/shipments/track/:trackingNumber", shipment.Track)

	// Start Server

	router.Use(middleware.Logger())
	for _, route := range router.Routes() {
		log.Printf("%s %s", route.Method, route.Path)
	}
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
