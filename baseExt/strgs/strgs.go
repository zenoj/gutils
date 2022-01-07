package strgs

// CountElements counts the number of occurrences of each number
// in the list l and adds the frequencies to the given map.
func CountElements(frequency map[string]int, l ...string) {
	for _, i := range l {
		frequency[i]++
	}
}

// DeleteElem deletes the ith element from the slice while changing the order of the slice
func DeleteElem(a []string, i int) {
	a[len(a)-1], a[i] = a[i], a[len(a)-1]
	a = a[:len(a)-1]
}

// DeleteElemKeepOrder deletes the ith element from the slice while preserving the previous order
func DeleteElemKeepOrder(a []string, i int) {
	a = append(a[:i], a[i+1:]...)
}
