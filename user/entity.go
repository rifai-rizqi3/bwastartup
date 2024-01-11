package user

import "time"

// User is a struct representing the user entity with various attributes.
type User struct {
	ID             int       // Unique identifier for the user.
	Name           string    // User's name.
	Occupation     string    // User's occupation.
	Email          string    // User's email address.
	PasswordHash   string    // Hashed password for security.
	AvatarFileName string    // File name of the user's avatar.
	Role           string    // Role or permission level of the user (e.g., "admin", "user").
	CreatedAt      time.Time // Timestamp indicating when the user was created.
	UpdatedAt      time.Time // Timestamp indicating when the user was last updated.
}
