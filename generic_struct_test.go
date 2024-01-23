package go_generics

import (
	"fmt"
	"testing"
)

type Product[T any] struct {
	Name T
	Id   T
}

func TestGenericStruct(t *testing.T) {
	data := Product[string]{
		Name: "Phone",
		Id:   "112233",
	}

	fmt.Println(data)
}
