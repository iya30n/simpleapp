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

func Validate(req *http.Request, validationRules Rule) error {
	for inputName, rule := range validationRules {
		reqInput := req.PostFormValue(inputName)

		for _, ruleName := range strings.Split(rule, "|") {
			if err := callValidator(inputName, ruleName, reqInput); err != nil {
				return err
			}
		}
	}

	return nil
}

var typesList []string = []string{"int", "int32", "int64", "float32", "float64", "string"}

func callValidator(inputName string, ruleName string, inputVal any) error {
	inputValType := fmt.Sprint(reflect.TypeOf(inputVal))

	// checking type of the input.
	if Array.Contains(ruleName, typesList) {
		if ruleName != inputValType {
			return fmt.Errorf("type of %s should be %s but %s given", inputName, ruleName, inputValType)
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
	fn := funcsList[ruleMethod]
	return fn(inputName, inputVal.(string), ruleValToInt)
}
