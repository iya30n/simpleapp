package validator

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"simpleapp/core/types/Array"
	"simpleapp/core/validator/contracts"
	stringrules "simpleapp/core/validator/rules/string"
	"strconv"
	"strings"
)

type Rule map[string]string

func Validate(req *http.Request, validationRules Rule) []error {
	var errors []error

	for inputName, rule := range validationRules {
		reqInput := req.PostFormValue(inputName)

		for _, ruleName := range strings.Split(rule, "|") {
			if ruleName == "required" {
				if reqInput == "" {
					errors = append(errors, fmt.Errorf("%s is required", inputName))
					break
				}

				continue
			}

			if err := callValidator(ruleName, reqInput); err != nil {
				errors = append(errors, fmt.Errorf("%s: %v", inputName, err.Error()))
			}
		}
	}

	return errors
}

// TODO: find a type or something to act as this.
var typesList []string = []string{"int", "int32", "int64", "float32", "float64", "string"}

func callValidator(ruleName string, inputVal any) error {
	inputValType := fmt.Sprint(reflect.TypeOf(inputVal))

	// checking type of the input.
	if Array.Contains(ruleName, typesList) {
		if ruleName != inputValType {
			return fmt.Errorf("should be %s but %s given", ruleName, inputValType)
		}

		return nil
	}

	// validation rule should be like: min:3
	splitRule := strings.Split(ruleName, ":")
	if len(splitRule) != 2 {
		log.Fatalf("the validation %s is not valid!", ruleName)
	}

	funcsList := map[string]contracts.ValidatorFunc{
		"min": stringrules.Min,
		"max": stringrules.Max,
	}

	// the value of ruleMethod = "min", ruleVal = 3
	ruleMethod, ruleVal := splitRule[0], splitRule[1]
	if _, ok := funcsList[ruleMethod]; !ok {
		log.Fatalf("the validation %s is not valid!", ruleMethod)
	}

	ruleValToInt, _ := strconv.Atoi(ruleVal)
	return funcsList[ruleMethod](inputVal.(string), ruleValToInt)
}
