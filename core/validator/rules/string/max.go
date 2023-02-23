package string

import "fmt"

func Max(value string, valLen int) error {
	if len(value) > valLen {
		return fmt.Errorf("should be less than %d characters", valLen)
	}

	return nil
}
