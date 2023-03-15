package others

import "fmt"

func Same(firstVal any, secondVal any) error {
	if firstVal != secondVal {
		return fmt.Errorf("%s is not the same as %s", firstVal.(string), secondVal.(string))
	}

	return nil
}
