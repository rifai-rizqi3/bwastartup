package user

import "gorm.io/gorm"

// Repository is an interface defining methods for user-related database operations.
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

// repository is a struct that implements the Repository interface.
type repository struct {
	db *gorm.DB
}

// NewRepository is a constructor function that creates a new repository instance.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save is a method of the repository that saves a user to the database.
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
