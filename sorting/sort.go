package main

// SelectionSortInt takes an unsorted []int array and preforms selection sort to sort the array passed in.
func selectionSortInt(n []int) {
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

func insertionSortInt(n []int) {
	for i := 1; i <= len(n)-1; i++ {
		if n[i] < n[i-1] {
			for j := i; j >= 1; j-- {
				// Loop back and compare all elements
				if n[j] < n[j-1] {
					// Swap
					n[j], n[j-1] = n[j-1], n[j]
				}
			}
		}
	}
}

// BubbleSortInt sorts an unsorted []int with the bubble sort algorithm
func bubbleSortInt(n []int) {
	sorted := false
	for !sorted {
		// True if we dont swap any
		sorted = true
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				// Swap means we arent sorted yet
				n[i], n[i+1] = n[i+1], n[i]
				sorted = false
			}
		}
	}
}
