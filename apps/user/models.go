package user

import (
	"gorm.io/gorm"

	app "github.com/stevo1403/go-by-example/initializers"
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

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) GetUserByID(userID int) User {
	var user User
	app.DB.Limit(1).First(&user, userID)

	return user
}

func (u User) Exists(userID int) bool {
	var user User
	result := app.DB.Limit(1).First(&user, userID)
	recordNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	return !recordNotFound

}
