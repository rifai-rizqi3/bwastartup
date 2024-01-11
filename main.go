package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Set up the database connection.
	dsn := "root:RNRif@i1212@tcp(localhost:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a user repository using GORM.
	userRepository := user.NewRepository(db)

	// Create a user service using the user repository.
	userService := user.NewService(userRepository)

	// Create a user handler using the user service.
	userHandler := handler.NewUserHandler(userService)

	// Create a Gin router.
	router := gin.Default()
	api := router.Group("/api/v1")

	// Set up a route for user registration.
	api.POST("/users", userHandler.RegisterUser)

	// Run the Gin router.
	router.Run()
}
