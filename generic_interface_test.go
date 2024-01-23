package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	data := MyData[string]{}

	result := ChangeValue(&data, "OK")

	assert.Equal(t, "OK", result)
	assert.Equal(t, "OK", data.Value)
}
