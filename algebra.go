package govector

import (
	"fmt"
	"math"
)

// Return a vector of element-wise products of two input vectors
func Product(x, y Vector) (Vector, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("x and y have unequal lengths: %d / %d", len(x), len(y))
	}

	p := make(Vector, len(x))
	for i, _ := range x {
		p[i] = x[i] * y[i]
	}
	return p, nil
}

// Return the dot product of two vectors
func DotProduct(x, y Vector) (float64, error) {
	p, err := Product(x, y)
	if err != nil {
		return 0, err
	}
	return p.Sum(), nil
}

// Returns the vector norm.  Use pow = 2.0 for Euclidean.
func Norm(x Vector, pow float64) float64 {
	s := 0.0

	for _, xval := range x {
		s += math.Pow(xval, pow)
	}

	return math.Pow(s, 1/pow)
}

// Return the cosine similarity between two vectors
func Cosine(x, y Vector) (float64, error) {
	d, err := DotProduct(x, y)
	if err != nil {
		return 0, err
	}

	xnorm := Norm(x, 2.0)
	ynorm := Norm(y, 2.0)

	return d / (xnorm * ynorm), nil
}

// Return the Pearson correlation between two vectors
func Cor(x, y Vector) (float64, error) {
	n := float64(len(x))
	xy, err := Product(x, y)
	if err != nil {
		return 0, err
	}

	sx, err := x.Sd()
	if err != nil {
		return 0, err
	}
	sy, err := y.Sd()
	if err != nil {
		return 0, err
	}

	mx, err := x.Mean()
	if err != nil {
		return 0, err
	}
	my, err := y.Mean()
	if err != nil {
		return 0, nil
	}

	r := (xy.Sum() - n*mx*my) / ((n - 1) * sx * sy)
	return r, nil
}
