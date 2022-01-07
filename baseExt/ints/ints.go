package ints

import (
	"strconv"
)

// ListToString converts an int array into a string array
func ListToString(l ...int) []string {
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
func DeleteElem(l []int, i int) {
	l[len(l)-1], l[i] = l[i], l[len(l)-1]
	l = l[:len(l)-1]
}

// DeleteElemKeepOrder deletes the ith element from the slice while preserving the previous order
func DeleteElemKeepOrder(l []int, i int) {
	l = append(l[:i], l[i+1:]...)
}

// ListContains checks if an integer e is in a integer list l
func ListContains(l []int, e int) bool {
	for _, i := range l {
		if e == i {
			return true
		}
	}
	return false
}
