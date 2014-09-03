# govector

Provide nice vector syntax for handling numeric types in [Go](http://golang.org).

## Usage

```go
# create a Vector type from an int array
x, err := AsVector([]int{1, 2, 3, 4, 6, 5})

# create a Vector type from a float64 array, to be used for weights
w, _ := AsVector([]float64{1.0, 1.0, 1.0, 1.0, 1.0, 4.0})

# find the differences of the Vector x
d_x, _ := x.Diff()

# Generate the empirical CDF function for x
empirical, _ := x.Ecdf()

# Calculate the percentile from the empirical CDF of x
percentile, _ = empirical(2.4)

# Calculate the weighted mean of x using weights w
wm, _ = x.WeightedMean(w)

# Calculate the 5% and 95% quantiles of x
q, _ := AsVector([]float64{0.05, 0.95})
quantiles, _ = x.Quantiles(q)

# cumulative sum
s, _ = x.Cumsum()

# shuffle x
shuffled, _ := x.Shuffle()

# Apply arbitrary functions to vectors
_ = x.Apply(empirical)
_ = x.Apply(math.Sqrt)
```
