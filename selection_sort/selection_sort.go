// Package selectionsort implements the classic selection sort algorithm.
package selectionsort

// SelectionSortInt takes an unsorted []int array and preforms selection sort .
func SelectionSortInt(n []int) {
	length := len(n)
	for i := 0; i < len(n); i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			// found a new minimum
			if n[j] < n[minIdx] {
				minIdx = j
			}
		}
		// swap new minimum with current index
		n[i], n[minIdx] = n[minIdx], n[i]
	}
}
