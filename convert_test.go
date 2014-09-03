package govector

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestConversion(t *testing.T) {
	x := []int64{2, 2, 2, 4, 2, 5}

	_, err := AsVector(x)
	assert.Equal(t, nil, err, "Error converting to vector type")

	_, err = AsVector(1)
	assert.NotEqual(t, nil, err, "Integer should return error in vector conversion")
}
