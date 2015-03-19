package govector

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestVectors(t *testing.T) {
	x, err := AsVector([]int{2, 2, 2, 4, 2, 5})
	assert.Equal(t, nil, err, "Error casting integer array to vector")

	w, err := AsVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})
	assert.Equal(t, nil, err, "Error casting float64 array to vector")

	q, err := AsVector([]float64{0.05, 0.95})
	assert.Equal(t, nil, err, "Error casing float64 array to vector")

	d_x := x.Diff()
	d_w := w.Diff()

	max := x.Max()
	assert.Equal(t, 5.0, max, "Error calculating max")

	min := x.Min()
	assert.Equal(t, 2.0, min, "Error calculating min")

	empirical := x.Ecdf()

	percentile := empirical(2.4)
	assert.Equal(t, 2.0/3.0, percentile, "Error in CDF calculation")

	_, err = d_x.WeightedMean(d_w)
	assert.Equal(t, nil, err, "Error calculating weighted mean")

	_ = x.Quantiles(q)

	cumsum := x.Cumsum()
	assert.Equal(t, Vector{2, 4, 6, 10, 12, 17}, cumsum, "Error calculating cumulative sum")

	ranks := x.Rank()
	assert.Equal(t, Vector{3, 0, 0, 4, 0, 5}, ranks, "Error calculating ranks")

	shuffled := x.Shuffle()
	assert.Equal(t, x.Len(), shuffled.Len(), "Error shuffling vector")

	y, err := AsVector([]int{-2, 2, -1, 4, 2, 5})
	assert.Equal(t, nil, err, "Error casting negative integer array to vector")

	abs := y.Abs()
	assert.Equal(t, Vector{2, 2, 1, 4, 2, 5}, abs, "Error finding absolute values")

	_ = x.Apply(empirical)

	n := x.Len()
	x.Push(50)
	assert.Equal(t, n+1, x.Len(), "Error appending value to vector")

	xw := Join(x, w)
	assert.Equal(t, x.Len()+w.Len(), xw.Len(), "Error joining vectors")

	filtered := xw.Filter(func(x float64) bool {
		if x < 10 {
			return false
		}
		return true
	})
	assert.Equal(t, 12, len(filtered), "Error filtering vector")

	x, _ = AsVector([]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18})
	smoothed, err := x.Smooth(0, 0)
	assert.Equal(t, nil, err, "Error smoothing vector")
	assert.Equal(t, x, smoothed, "For i and j equal to zero, we'd expect to return the input vector after smoothing")
	smoothed, err = x.Smooth(1, 1)
	expected := Vector{1, 2, 4, 6, 8, 10, 12, 14, 16, 17}
	assert.Equal(t, nil, err, "Error smoothing vector")
	assert.Equal(t, expected, smoothed, "Error smoothing vector")
	smoothed, err = x.Smooth(-1, -1)
	assert.NotEqual(t, nil, err, "Negative input indices should return an error")
	assert.Equal(t, x, smoothed, "Negative input indices should return the original vector")
}
