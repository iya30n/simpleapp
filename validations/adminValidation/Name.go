package adminValidation

import "fmt"

func ValidateName(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("Name is too short!")
	}

	return nil
}