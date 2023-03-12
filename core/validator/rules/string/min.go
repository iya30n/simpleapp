package string

import "fmt"

func Min(value any, valLen any) error {
	if len(value.(string)) < valLen.(int) {
		return fmt.Errorf("should be greater than %d characters", valLen)
	}

	return nil
}
