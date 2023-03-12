package others

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameFunc(t *testing.T) {
	firstVal, secondVal := "test", "test"

	sameFuncErr := Same(firstVal, secondVal)

	assert.NoError(t, sameFuncErr)
}

func TestSameFuncWithNotEqualValues(t *testing.T) {
	firstVal, secondVal := "first", "second"

	errMsg := fmt.Sprintf("%s is not the same as %s", firstVal, secondVal)

	sameFuncErr := Same(firstVal, secondVal)

	assert.EqualError(t, sameFuncErr, errMsg)
}