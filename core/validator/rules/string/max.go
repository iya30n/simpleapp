package string

import "fmt"

func Max(inputName string, value string, valLen int) error {
	if len(value) > valLen {
		return fmt.Errorf("the length of %s should be less than %d", inputName, valLen)
	}

	return nil
}
