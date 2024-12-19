package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/stevo1403/go-by-example/initializers"
	"gorm.io/gorm"
)

type UserView interface {
	GetUser()
	GetUsers()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

func CreateUser(c *gin.Context) {
	var MIN_PASSWORD_LENGTH = 7

	var userBody UserSchema

	// Convert request body to schema
	err := c.BindJSON(&userBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert schema to an Object
	userObj := userBody.to_object()

	// Check password length
	if len(userObj.Password) < MIN_PASSWORD_LENGTH {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Password must be at least %d characters long.", MIN_PASSWORD_LENGTH)},
		)
		return
	}
	// Check phone number validity
	if !isValidPhoneNumber(userObj.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number."})
		return
	}
	// Check email validity
	if !isValidEmailRegex(userObj.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address."})
		return
	}
	// Is email registered to another account?
	emailAlreadyExists := userObj.EmailExists(userObj.Email)
	if emailAlreadyExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email address is already in use. Please choose a different email address."})
		return
	}
	// Generate password
	_, err = userObj.GeneratePasswordHash()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Create user in the DB
	err = userObj.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	respObj := UserOut{
		User: userObj.to_schema(),
	}

	c.JSON(http.StatusOK, gin.H{"data": respObj})
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	var user User
	// Get the user by ID
	app.DB.Limit(1).First(&user, userId)

	// Convert the user to schema
	respObj := UserOut{
		User: user.to_schema(),
	}

	c.JSON(http.StatusOK, gin.H{"data": respObj})
}

func GetUsers(c *gin.Context) {
	var users []User
	app.DB.Find(&users)

	var users_as_schema []UserOutSchema
	for _, user := range users {
		users_as_schema = append(users_as_schema, user.to_schema())
	}

	respObj := UsersOut{Users: users_as_schema}
	c.JSON(http.StatusOK, gin.H{"data": respObj})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	userBody := UserSchema{}
	c.BindJSON(&userBody)

	var user User
	result := app.DB.Limit(1).First(&user, userId)

	if result.Error == gorm.ErrRecordNotFound || result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data":    map[string]interface{}{},
			"message": fmt.Sprintf("User identified by user ID '%s' does not exist", userId),
		})
	} else {
		// Update the user
		user.FirstName = userBody.FirstName
		user.LastName = userBody.LastName
		user.Email = userBody.Email
		user.Phone = userBody.Phone

		app.DB.Save(&user)

		respObj := user.to_schema()
		c.JSON(http.StatusOK, gin.H{"data": respObj})
	}

}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	var userBody UserSchema

	c.BindJSON(&userBody)

	var user User
	err := app.DB.Limit(1).First(&user, userId).Error

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"data":    map[string]interface{}{},
			"message": fmt.Sprintf("User identified by user ID '%s' does not exist", userId),
		})
	} else {

		// Delete the user
		result := app.DB.Delete(&user)
		resultNotDeleted := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

		if resultNotDeleted {
			c.JSON(http.StatusNotFound, gin.H{
				"data":    map[string]interface{}{},
				"message": fmt.Sprintf("An error occurred: User with user id '%s' could not be deleted.", userId),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    map[string]interface{}{},
			"message": "User deleted successfully",
		})
	}

}

func AuthenticateUser(c *gin.Context) {
	var userBody UserSchema
	c.BindJSON(&userBody)

	var user User
	result := app.DB.Limit(1).Where(&User{Email: userBody.Email}).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"data":    map[string]interface{}{},
			"message": "User not found",
		})
		return
	}

	// Check password
	passwordIsCorrect := user.ComparePassword(userBody.Password)
	if !passwordIsCorrect {
		c.JSON(http.StatusUnauthorized, gin.H{
			"data":    map[string]interface{}{},
			"message": "Invalid password",
		})
		return
	}

	user_token, err := user.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    map[string]interface{}{},
			"message": "An error occurred while generating token",
		})
		return
	}

	respObj := LoginOut{
		User:  user.to_schema(),
		Token: user_token,
	}

	c.JSON(http.StatusOK, gin.H{"data": respObj})
}
