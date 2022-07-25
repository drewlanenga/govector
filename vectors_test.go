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

	m, v := Vector{1., 2., 3.}.MeanVar()
	assert.Equal(t, 2., m, "Incorrect mean calculation")
	assert.Equal(t, 1., v, "Incorrect variance calculation")

	_, err = d_x.WeightedMean(d_w)
	assert.Equal(t, nil, err, "Error calculating weighted mean")

	_ = x.Quantiles(q)

	cumsum := x.Cumsum()
	assert.Equal(t, Vector{2, 4, 6, 10, 12, 17}, cumsum, "Error calculating cumulative sum")

	ranks := x.Rank()
	assert.Equal(t, Vector{0, 0, 0, 4, 0, 5}, ranks, "Error calculating ranks")

	order := x.Order()
	assert.Equal(t, Vector{0, 1, 2, 4, 3, 5}, order, "Error calculating order")

	shuffled := x.Shuffle()
	assert.Equal(t, x.Len(), shuffled.Len(), "Error shuffling vector")

	unique := x.Unique()
	assert.Equal(t, Vector{2.0, 4.0, 5.0}, unique, "Error getting unique values")

	sub := Vector{1.0, 0.0, 0.5, 3.5, 0.2, 0.3}
	difference, err := x.Subtract(sub)
	assert.Equal(t, nil, err, "Differing lengths in subtracting two vectors")
	assert.Equal(t, Vector{1.0, 2.0, 1.5, 0.5, 1.8, 4.7}, difference, "Error subtracting vectors")

	cdifference := x.SubtractConst(1.0)
	assert.Equal(t, Vector{1.0, 1.0, 1.0, 3.0, 1.0, 4.0}, cdifference, "Error subtracting constant from vector")

	toround := Vector{0.1134224, 0.29985, 0.00081, 0.2, 0.5555}
	rounded := toround.Round(2)
	assert.Equal(t, Vector{0.11, 0.30, 0.0, 0.2, 0.56}, rounded, "Error rounding vector")

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

	z, err := AsVector([]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18})
	assert.Equal(t, nil, err)

	smoothed := z.Smooth(0, 0)
	assert.Equal(t, z, smoothed)

	smoothed = z.Smooth(1, 1)
	expected := Vector{1, 2, 4, 6, 8, 10, 12, 14, 16, 17}
	assert.Equal(t, expected, smoothed, "Error smoothing vector")

	x.Sort()
	assert.Equal(t, Vector{2, 2, 2, 2, 4, 5, 50}, x)
}

func TestFixedPush(t *testing.T) {
	arr := make([]float64, 3, 3)

	v := Vector(arr)
	err := v.PushFixed(5.0)
	err = v.PushFixed(25.0)
	err = v.PushFixed(125.0)
	assert.Equal(t, v[2], 125.0)

	err = v.PushFixed(250.0)
	err = v.PushFixed(350.0)
	assert.Equal(t, err, nil)
	assert.Equal(t, v[2], 350.0)
	assert.Equal(t, v[0], 125.0)
	assert.Equal(t, len(v), 3)
}
