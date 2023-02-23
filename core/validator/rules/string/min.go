package string

import "fmt"

func Min(value string, valLen int) error {
	if len(value) < valLen {
		return fmt.Errorf("should be greater than %d characters", valLen)
	}

	return nil
}
