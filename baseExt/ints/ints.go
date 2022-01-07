package ints

import (
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

// CountElements counts the number of occurrences of each number
// in the list l and adds the frequencies to the given map.
func CountElements(frequency map[int]int, l ...int) {
	for _, i := range l {
		frequency[i]++
	}
}
