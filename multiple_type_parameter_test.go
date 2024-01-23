package go_generics

import (
	"fmt"
	"testing"
)

func MultipleParameter[T any, X any](param1 T, param2 X) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Hello World", 200)
}
