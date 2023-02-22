package validator

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"simpleapp/core/types/Array"
	"strings"
)

type Rule map[string]string

func Validate(req http.Request, validationRules Rule) error {
	for inputName, rule := range validationRules {
		reqInput := req.PostFormValue(inputName)

		for _, ruleName := range strings.Split(rule, "|") {
			callValidator(inputName, ruleName, reqInput)
		}
	}

	return nil
}

type dType interface {
	string | bool | int | int32 | int64 | float32 | float64
}

func callValidator[dtype dType](inputName string, ruleName string, inputVal dtype) error {
	typesToCheck := []string{"string", "bool", "int", "int32", "int64", "float32", "float64"}

	inputValType := fmt.Sprint(reflect.TypeOf(inputVal))
	// checking type of the input.
	if Array.Contains(ruleName, typesToCheck) && ruleName != inputValType {
		return fmt.Errorf("type of %s should be %s but %s given", inputName, ruleName, inputValType)
	}

	// validation rule should be like: min:3
	splitRule := strings.Split(ruleName, ":")
	if len(splitRule) != 2 {
		log.Fatalf("the validation %s is not valid!", ruleName)
	}

	// the value of ruleMethod = "min", ruleVal = 3
	ruleMethod, ruleVal := splitRule[0], splitRule[1]
	// TODO: call the ruleMethod from core/validator/validations and return the result
	
	/* f := reflect.ValueOf(ruleMethod)
	result := f.Call([]reflect.Value{reflect.ValueOf(ruleName), reflect.ValueOf(inputVal), reflect.ValueOf(ruleVal)})
	return result[0]. */

	return nil
}
