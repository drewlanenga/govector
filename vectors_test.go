package govector

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestVectors(t *testing.T) {
	x, err := AsVector([]int{2, 2, 2, 4, 2, 5})
	assert.Equal(t, nil, err, "Error casting integer array to vector")

	w, err := AsVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})
	assert.Equal(t, nil, err, "Error casting float64 array to vector")

	q, err := AsVector([]float64{0.05, 0.95})
	assert.Equal(t, nil, err, "Error casing float64 array to vector")

	d_x, err := x.Diff()
	assert.Equal(t, nil, err, "Error calculating vector differences")

	d_w, err := w.Diff()
	assert.Equal(t, nil, err, "Error calculating vector differences")

	max := x.Max()
	assert.Equal(t, 5.0, max, "Error calculating max")

	min := x.Min()
	assert.Equal(t, 2.0, min, "Error calculating min")

	empirical, err := x.Ecdf()
	assert.Equal(t, nil, err, "Error creating CDF function")

	_ = empirical(2.4)

	_, err = d_x.WeightedMean(d_w)
	assert.Equal(t, nil, err, "Error calculating weighted mean")

	_, err = x.Quantiles(q)
	assert.Equal(t, nil, err, "Error calculating quantiles")

	_, err = x.Cumsum()
	assert.Equal(t, nil, err, "Error calculating cumulative sum")

	_, err = x.Rank()
	assert.Equal(t, nil, err, "Error calculating ranks")

	_, err = x.Shuffle()
	assert.Equal(t, nil, err, "Error shuffling vector")

	y, err := AsVector([]int{-2, 2, -1, 4, 2, 5})
	assert.Equal(t, nil, err, "Error casting negative integer array to vector")

	_, err = y.Abs()
	assert.Equal(t, nil, err, "Error finding absolute values")

	_ = x.Apply(empirical)
}
