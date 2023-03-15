package Array

func Contains[dType string|int](value dType, list []dType) bool {
	for _, val := range list {
		if value == val {
			return true
		}
	}

	return false
}