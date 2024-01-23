package go_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParameter(t *testing.T) {
	result := Length[string]("Hello Type Parameter String")
	//fmt.Println(result)
	assert.Equal(t, "Hello Type Parameter String", result)

	resultNumber := Length[int](100)
	//fmt.Println(resultInt)
	assert.Equal(t, 100, resultNumber)
}
