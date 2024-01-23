package go_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestGenericType(t *testing.T) {
	resultNumbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(resultNumbers)

	resultString := Bag[string]{"Jakarta", "Bandung", "Jogja"}
	PrintBag(resultString)
}
