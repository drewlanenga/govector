package govector

import (
	"fmt"
	"math"
)

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

func DotProduct(x, y Vector) (float64, error) {
	p, err := Product(x, y)
	if err != nil {
		return 0, err
	}
	return p.Sum()
}

// Returns the vector norm.  Use pow = 2.0 for Euclidean.
func Norm(x Vector, pow float64) float64 {
	s := 0.0

	for _, xval := range x {
		s += math.Pow(xval, pow)
	}

	return math.Pow(s, 1/pow)
}

func Cosine(x, y Vector) (float64, error) {
	d, err := DotProduct(x, y)
	if err != nil {
		return 0, err
	}

	xnorm := Norm(x, 2.0)
	ynorm := Norm(y, 2.0)

	return d / (xnorm * ynorm), nil
}
