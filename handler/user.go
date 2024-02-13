package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// userHandler is a struct that holds the user service.
type userHandler struct {
	userService user.Service
}

// NewUserHandler is a constructor function that creates a new userHandler instance.
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// RegisterUser is a method of the userHandler that handles user registration.
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	// Try to bind JSON data from the request body into the input struct.
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Attempt to register the user using the user service.
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// If registration fails, respond with an error.
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// TODO: Implement a proper token generation mechanism
	// For now, a placeholder token is used.
	token := "tokentokentokentoken"

	// Format the user and token for the response.
	formatter := user.FormatUser(newUser, token)

	// Respond with a success message and user information.
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Login Successfuly", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
