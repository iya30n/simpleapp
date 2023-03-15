package Array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsFuncWithString(t *testing.T) {
	arr := []string{"first", "second", "third"}

	assert.Equal(t, true, Contains("second", arr))
	assert.Equal(t, false, Contains("mamad", arr))
}

func TestContainsFuncWithInt(t *testing.T) {
	arr := []int{1, 2, 5, 7}

	assert.Equal(t, true, Contains(2, arr))
	assert.Equal(t, false, Contains(3, arr))
}

