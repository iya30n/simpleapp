package string

import (
	"fmt"
	"strconv"
)

func Min(value any, valLen any) error {
	vLen, _ := strconv.Atoi(valLen.(string))

	if len(value.(string)) < vLen {
		return fmt.Errorf("should be greater than %d characters", vLen)
	}

	return nil
}
