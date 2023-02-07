package login

import "fmt"

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password should be more than 8 characters")
	}

	if len(password) > 100 {
		return fmt.Errorf("password should be less than 100 characters")
	}

	return nil
}