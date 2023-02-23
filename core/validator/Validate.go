package validator

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"simpleapp/core/validator/contracts"
	"simpleapp/core/validator/rules/string"
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

func callValidator(inputName string, ruleName string, inputVal any) error {
	inputValType := fmt.Sprint(reflect.TypeOf(inputVal))
	// checking type of the input.
	if ruleName == inputValType && ruleName != inputValType {
		return fmt.Errorf("type of %s should be %s but %s given", inputName, ruleName, inputValType)
	}

	// validation rule should be like: min:3
	splitRule := strings.Split(ruleName, ":")
	if len(splitRule) != 2 {
		log.Fatalf("the validation %s is not valid!", ruleName)
	}

	funcsList := map[string]contracts.ValidatorFunc{
		"min": string.Min,
		"max": string.Max,
	}

	// the value of ruleMethod = "min", ruleVal = 3
	ruleMethod, ruleVal := splitRule[0], splitRule[1]
	if _, ok := funcsList[ruleMethod]; !ok {
		log.Fatalf("the validation %s is not valid!", ruleMethod)
	}

	ruleValToInt, _ := strconv.Atoi(ruleVal)

	return funcsList[ruleMethod](inputName, inputVal.(string), ruleValToInt)
}
