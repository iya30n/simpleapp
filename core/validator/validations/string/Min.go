package string

import "fmt"

func Min(inputName string, value string, valLen int) error {
	if len(value) < valLen {
		return fmt.Errorf("the length of %s should be greater than %d", inputName, valLen)
	}

	return nil
}
