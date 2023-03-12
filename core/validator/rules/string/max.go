package string

import "fmt"

func Max(value any, valLen any) error {
	if len(value.(string)) > valLen.(int) {
		return fmt.Errorf("should be less than %d characters", valLen)
	}

	return nil
}
