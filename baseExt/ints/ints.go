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

// DeleteElem deletes the ith element from the slice while changing the order of the slice
func DeleteElem(a []int, i int) {
	a[len(a)-1], a[i] = a[i], a[len(a)-1]
	a = a[:len(a)-1]
}

// DeleteElemKeepOrder deletes the ith element from the slice while preserving the previous order
func DeleteElemKeepOrder(a []int, i int) {
	a = append(a[:i], a[i+1:]...)
}
