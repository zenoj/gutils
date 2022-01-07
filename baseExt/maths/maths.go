package maths

import "math"

// Abs returns the absolut value of a given integers
func Abs(e int) int {
	if e < 0 {
		e = -e
	}
	return e
}

// MaxInt returns the maximum of multiple integers
func MaxInt(e ...int) int {
	max := math.MinInt
	for _, i := range e {
		if i > max {
			max = i
		}
	}
	return max
}

// MinInt returns the minimum of multiple integers
func MinInt(e ...int) int {
	min := math.MinInt
	for _, i := range e {
		if i < min {
			min = i
		}
	}
	return min
}

// Sum computes the sum of multiple integers
func Sum(e ...int) int {
	res := 0
	for _, i := range e {
		res += i
	}
	return res
}

// Product computes the product of multiple integers
func Product(e ...int) int {
	res := 1
	for _, i := range e {
		res *= i
	}
	return res
}
