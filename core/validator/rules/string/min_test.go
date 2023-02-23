package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFunc(t *testing.T) {
	errMsg := fmt.Sprintf("should be greater than %d characters", 3)

	assert.EqualError(t, Min("si", 3), errMsg)

	assert.NoError(t, Min("sina21", 3))
}
