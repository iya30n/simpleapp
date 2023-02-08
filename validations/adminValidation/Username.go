package adminValidation

import "fmt"

func ValidateUsername(username string) error {
	if len(username) < 3 {
		return fmt.Errorf("username should be more than 3 characters")
	}

	if len(username) > 100 {
		return fmt.Errorf("username should be less than 100 characters")
	}

	return nil
}