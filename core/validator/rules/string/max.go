package string

import (
	"fmt"
	"strconv"
)

func Max(value any, valLen any) error {
	vLen, _ := strconv.Atoi(valLen.(string))

	if len(value.(string)) > vLen {
		return fmt.Errorf("should be less than %d characters", vLen)
	}

	return nil
}
