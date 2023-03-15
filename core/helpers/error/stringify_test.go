package error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringifyErrorsList(t *testing.T) {
	errList := []error{
		fmt.Errorf("first error"),
		fmt.Errorf("second error"),
		fmt.Errorf("third error"),
	}

	stringErrors := Stringify(errList)

	assert.Equal(t, stringErrors, []string{"first error", "second error", "third error"})
}
