package user

import "golang.org/x/crypto/bcrypt"

// Service is an interface defining methods for user-related operations.
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

// service is a struct that implements the Service interface.
type service struct {
	repository Repository
}

// NewService is a constructor function that creates a new service instance.
func NewService(repository Repository) *service {
	return &service{repository}
}

// RegisterUser is a method of the service that registers a new user.
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	// Hash the user's password using bcrypt.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	// Save the new user to the repository.
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
