package user

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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

func (u User) GetUserByID(userID uint) User {
	var user User
	app.DB.Limit(1).First(&user, userID)

	return user
}

func (u User) Exists(userID uint) bool {
	var user User
	result := app.DB.Limit(1).First(&user, userID)
	recordNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	return !recordNotFound

}

func (u User) EmailExists(email string) bool {
	var user User
	result := app.DB.Limit(1).Where(&User{Email: email}).First(&user)

	recordNotFound := (result.Error != nil || result.Error == gorm.ErrRecordNotFound)

	return !recordNotFound
}

func (u *User) UpdatePasswordHash(password string) error {
	// Check if user exists
	if !u.Exists(u.ID) {
		return fmt.Errorf("cannot update password for non-existing user")
	}

	// Convert plaintext password to bytes
	password_bytes := []byte(password)

	// Hash the password
	hashed_password, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)

	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return fmt.Errorf("cannot generate password: password is too long")
		} else {
			return fmt.Errorf("cannot generate password: %w", err)
		}
	}
	// Convert the hashed password to a string
	password_s := string(hashed_password)

	// Set the user password to the hashed password
	u.Password = password_s

	return nil

}

func (u *User) GeneratePasswordHash() (string, error) {

	if u.Exists(u.ID) {
		return "", fmt.Errorf("cannot generate password for existing user")
	}
	// Convert plaintext password to bytes
	password_bytes := []byte(u.Password)

	// Hash the password
	hashed_password, err := bcrypt.GenerateFromPassword(password_bytes, bcrypt.DefaultCost)

	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return "", fmt.Errorf("cannot generate password: password is too long")
		} else {
			return "", fmt.Errorf("cannot generate password: %w", err)
		}
	}
	// Convert the hashed password to a string
	password := string(hashed_password)

	// Set the user password to the hashed password
	u.Password = password

	return password, nil
}

func (u User) ComparePassword(inputPassword string) bool {
	// Check if password exists.
	if u.Password == "" || inputPassword == "" {
		return false
	}

	// Compare user-supplied password against DB-stored password
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))

	if err == nil {
		return true
	} else {
		return false
	}
}

func (u *User) Create() error {
	result := app.DB.Create(&u)

	if result.Error != nil {
		return fmt.Errorf("cannot create user: %w", result.Error)
	}

	return nil

}

func (u User) GenerateToken() (string, error) {
	if !u.Exists(u.ID) {
		return "", fmt.Errorf("cannot generate token: user doesn't exist")
	} else {
		// Get the HMAC signing key
		auth_token_key := os.Getenv("AUTH_TOKEN_KEY")

		expirationDuration := time.Hour * 2

		token := jwt.New(
			jwt.SigningMethodHS256,
		)

		standard_claims := &jwt.StandardClaims{
			Issuer:    "gbe",
			Subject:   strconv.Itoa(int(u.ID)),
			Audience:  "users",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expirationDuration).Unix(),
			NotBefore: time.Now().Unix(),
		}

		token.Claims = &UserJWTClaim{
			StandardClaims: *standard_claims,
			Email:          u.Email,
			Roles:          make([]string, 0),
		}

		tokenString, err := token.SignedString([]byte(auth_token_key))

		if err != nil {
			return "", fmt.Errorf("cannot generate token: %w", err)
		} else {
			return tokenString, nil
		}
	}
}

func (u User) VerifyToken(tokenString string) (bool, error) {
	// Get the HMAC signing key
	auth_token_key := os.Getenv("AUTH_TOKEN_KEY")
	auth_token_key_bytes := []byte(auth_token_key)

	AUDIENCE := "users"
	ISSUER := "gbe"

	current_time := time.Now().Unix()
	claims := &UserJWTClaim{}

	_token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) { return auth_token_key_bytes, nil },
	)

	if current_time > claims.ExpiresAt {
		return false, fmt.Errorf("token verification failed: %s", "token is expired")
	} else if current_time < claims.NotBefore {
		return false, fmt.Errorf("token verification failed: %s", "token cannot be used before the stated time")
	} else if current_time < claims.IssuedAt {
		return false, fmt.Errorf("token verification failed: %s", "token was issued in the future")
	} else if claims.Audience != AUDIENCE {
		return false, fmt.Errorf("token verification failed: %s", "token belongs to a difference audience")
	} else if claims.Issuer != ISSUER {
		return false, fmt.Errorf("token verification failed: %s", "token wasn't issued by a recognized party")
	} else if !_token.Valid || err != nil {
		return false, fmt.Errorf("token verification failed: %w", err)
	}

	return true, nil

}

type UserJWTClaim struct {
	jwt.StandardClaims
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
