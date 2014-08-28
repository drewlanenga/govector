package govector

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAlgebra(t *testing.T) {
	x := IntToVector([]int{1, 2, 3, 4, 6, 5})
	y := IntToVector([]int{2, 1, 3, 4, 5, 6})

	var err error
	_, err = Product(x, y)
	assert.Equal(t, nil, err, "Error calculating vector product")

	_, err = DotProduct(x, y)
	assert.Equal(t, nil, err, "Error calculating dot product")

	_, err = Cosine(x, y)
	assert.Equal(t, nil, err, "Error calculating cosine similarity")

	_, err = Cor(x, y)
	assert.Equal(t, nil, err, "Error calculating vector correlation")
}
