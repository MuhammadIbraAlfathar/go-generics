package go_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Product[T any] struct {
	Name T
	Id   T
}

func (p *Product[T]) SayHello(name string) string {
	return "Hello " + name
}

func (p *Product[T]) ChangeName(name T) T {
	p.Name = name
	return p.Name
}

func TestGenericStruct(t *testing.T) {
	data := Product[string]{
		Name: "Phone",
		Id:   "112233",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Product[string]{
		Name: "Phone",
		Id:   "112233",
	}

	assert.Equal(t, "Fashion", data.ChangeName("Fashion"))
	assert.Equal(t, "Hello Ibra", data.SayHello("Ibra"))
}
