package user

import (
	"regexp"

	"github.com/badoux/checkmail"
	"github.com/nyaruka/phonenumbers"
)

func isValidEmail(email string) bool {
	err := checkmail.ValidateFormat(email)

	return err == nil
}

func isValidEmailRegex(email string) bool {
	// Email validation Regex
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func isValidPhoneNumber(phone string) bool {
	// Phone number validation Regex
	phoneNumber, err := phonenumbers.Parse(phone, phonenumbers.UNKNOWN_REGION)

	if err != nil {
		return false
	} else {
		return phonenumbers.IsValidNumber(phoneNumber)
	}
}
