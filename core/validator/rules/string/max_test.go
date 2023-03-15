package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFunc(t *testing.T) {
	errMsg := fmt.Sprintf("should be less than %d characters", 3)

	assert.EqualError(t, Max("alsfjlsdkfjlsdkfj", "3"), errMsg)

	assert.NoError(t, Max("sin", "3"))
}
