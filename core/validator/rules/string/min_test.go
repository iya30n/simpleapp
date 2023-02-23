package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFunc(t *testing.T) {
	errMsg := fmt.Sprintf("the length of %s should be greater than %d", "username", 3)

	assert.EqualError(t, Min("username", "si", 3), errMsg)

	assert.NoError(t, Min("username", "sina21", 3))
}
