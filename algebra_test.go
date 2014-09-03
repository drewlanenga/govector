package govector

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAlgebra(t *testing.T) {
	x, err := AsVector([]int{1, 2, 3, 4, 6, 5})
	assert.Equal(t, nil, err, "Error casting integer array to vector")

	y, err := AsVector([]int{2, 1, 3, 4, 5, 6})
	assert.Equal(t, nil, err, "Error casting integer array to vector")

	_, err = Product(x, y)
	assert.Equal(t, nil, err, "Error calculating vector product")

	_, err = DotProduct(x, y)
	assert.Equal(t, nil, err, "Error calculating dot product")

	_, err = Cosine(x, y)
	assert.Equal(t, nil, err, "Error calculating cosine similarity")

	_, err = Cor(x, y)
	assert.Equal(t, nil, err, "Error calculating vector correlation")
}
