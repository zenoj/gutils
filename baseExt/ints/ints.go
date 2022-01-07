package ints

import (
	"math"
	"strconv"
)

// IntsToString converts an int array into a string array
func IntsToString(l ...int) []string {
	a := make([]string, 0, len(l))
	for _, e := range l {
		a = append(a, strconv.Itoa(e))
	}
	return a
}

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

// CountElements counts the number of occurrences of each number
// in the list l and adds the frequencies to the given map.
func CountElements(frequency map[int]int, l ...int) {
	for _, i := range l {
		frequency[i]++
	}
}
