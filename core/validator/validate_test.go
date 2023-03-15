package validator

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	req := &http.Request{
		PostForm: make(map[string][]string),
	}

	validationRules := map[string]string{
		"username": "required|min:3|max:50",
		"password": "required|min:3|max:10",
	}

	errorsExpected := []string{"username is required", "password is required"}
	validateResult := Validate(req, validationRules)

	assert.Len(t, validateResult, len(errorsExpected))

	for i, val := range errorsExpected {
		assert.EqualError(t, validateResult[i], val)
	}
}

func TestValidateMinMax(t *testing.T) {
	req := &http.Request{
		PostForm: make(map[string][]string),
	}

	req.PostForm.Add("username", "si")
	req.PostForm.Add("password", "sia;sdfjalsdjf;asdjflsdjflajdkfalskdfjlasdjflkajsdlkfjasdjf")

	validationRules := map[string]string{
		"username": "required|min:3|max:50",
		"password": "required|min:3|max:10",
	}

	errorsExpected := []string{"username: should be greater than 3 characters", "password: should be less than 10 characters"}
	validateResult := Validate(req, validationRules)

	assert.Len(t, validateResult, len(errorsExpected))

	for i, val := range errorsExpected {
		assert.EqualError(t, validateResult[i], val)
	}
}

func TestValidateSame(t *testing.T) {
	req := &http.Request{
		PostForm: make(map[string][]string),
	}

	req.PostForm.Add("username", "sina")

	validationRules := map[string]string{
		"username": "required|min:3|max:50|same:mamad",
	}

	errorsExpected := []string{"username: sina is not the same as mamad"}
	validateResult := Validate(req, validationRules)

	assert.Len(t, validateResult, len(errorsExpected))

	for i, val := range errorsExpected {
		assert.EqualError(t, validateResult[i], val)
	}
}
