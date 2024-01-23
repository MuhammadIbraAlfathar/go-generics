package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type Number interface {
	~int | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](value1, value2 T) T {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}
