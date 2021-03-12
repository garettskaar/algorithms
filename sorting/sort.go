package main

// selectionSortInt takes an unsorted []int array and preforms selection sort.
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

// insertionSortInt takes an unsorted []int array and preforms insertion sort.
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

// bubbleSortInt takes an unsorted []int array and preforms bubble sort.
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

// quickSortInt takes an unsorted []int array and preforms quick sort.
func quickSortInt(n []int, l int, h int) {
	if l < h {

		pi := qsPartition(n, l, h)

		quickSortInt(n, l, pi-1)
		quickSortInt(n, pi+1, h)
	}
}

// partition for quicksort - uses high index as pivot
func qsPartition(n []int, l int, h int) int {
	// Pivot
	p := n[h]
	// Smaller element index
	i := l - 1

	for j := l; j <= h-1; j++ {
		// Smaller than pivot
		if n[j] < p {
			// increment small element index
			i++
			// Swap
			n[i], n[j] = n[j], n[i]
		}
	}
	// Swap pivot with correct position based on index i (greatest of all smaller elements and least of all bigger elements)
	n[i+1], n[h] = n[h], n[i+1]
	return i + 1
}
func mergeSortInt(n []int, l int, r int) {
	if l < r {
		// Middle division point
		m := l + (r-l)/2
		mergeSortInt(n, l, m)
		mergeSortInt(n, m+1, r)
		merge(n, l, m, r)
	}
}

// merge for merge sort. Combine subarrays into final merged result
func merge(n []int, l int, m int, r int) {
	// Lengths of sub arrays
	n1 := m - l + 1
	n2 := r - m
	// New temp sub arrays
	left := make([]int, n1)
	right := make([]int, n2)
	// Init left
	for i := 0; i < n1; i++ {
		left[i] = n[l+i]
	}
	// Init right
	for j := 0; j < n2; j++ {
		right[j] = n[m+1+j]
	}
	// initial indexes
	i := 0
	j := 0
	k := l
	// Merge
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			n[k] = left[i]
			i++
		} else {
			n[k] = right[j]
			j++
		}
		k++
	}
	// Pick up any remaining on the left and right
	for i < n1 {
		n[k] = left[i]
		i++
		k++
	}
	for j < n2 {
		n[k] = right[j]
		j++
		k++
	}
}
