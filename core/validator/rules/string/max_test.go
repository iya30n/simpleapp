package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFunc(t *testing.T) {
	errMsg := fmt.Sprintf("the length of %s should be less than %d", "username", 3)

	assert.EqualError(t, Max("username", "alsfjlsdkfjlsdkfj", 3), errMsg)

	assert.NoError(t, Max("username", "sin", 3))
}
