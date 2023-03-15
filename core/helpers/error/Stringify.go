package error

func Stringify(errors []error) []string {
	var strErrors []string

	for _, e := range errors {
		strErrors = append(strErrors, e.Error())
	}

	return strErrors
}