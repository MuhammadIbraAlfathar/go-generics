package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	result := IsSame[string]("hello", "hello")
	resultNumber := IsSame[int](200, 200)

	assert.Equal(t, true, result)
	assert.Equal(t, true, resultNumber)
}
