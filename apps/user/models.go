package user

import (
	"gorm.io/gorm"
)

// Create model for the user

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}
