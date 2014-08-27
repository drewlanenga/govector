package govector

import ()

//
//  there has got to be a better way to do these type conversions
//	... and ignore complex types for now
//

func IntToVector(x []int) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Uint8ToVector(x []uint8) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Uint16ToVector(x []uint16) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Uint32ToVector(x []uint32) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Uint64ToVector(x []uint64) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Int8ToVector(x []int) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Int16ToVector(x []int16) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Int32ToVector(x []int32) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Int64ToVector(x []int64) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Float32ToVector(x []float32) Vector {
	y := make(Vector, len(x))

	for i, _ := range x {
		y[i] = float64(x[i])
	}
	return y
}

func Float64ToVector(x []float64) Vector {
	return Vector(x)
}
