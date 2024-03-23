package main

import "fmt"

// binarySearch searches for a target string in a sorted list of string using binary search
// It returns the index of the target string if found, otherwise returns -1.
func binarySearch(names []string, target string) int {
	// Initialize the lower bound of the search range
	low := 0

	// Initialize the upper bound of the search range.
	high := len(names) - 1

	// Perform binary search until the search range is valid.
	for low <= high {
		// Calculate the middle index of the search range.
		mid := (low + high) / 2

		// If the middle element is equal to the target string, return its index.
		if names[mid] == target {
			return mid
		} else if names[mid] < target {
			// If the middle element is less than target string, adjust the lower bound.
			low = mid + 1
		} else {
			// If the middle element is greater than the target string, adjust the upper bound.
			high = mid - 1
		}
	}

	// Return -1 if the tartget string is not found
	return -1
}

func main() {
	names := []string{
		"Alice",
		"Bob",
		"Charlie",
		"David",
		"Emma",
		"Frank",
		"Gorge",
	}

	target := "Frank"

	index := binarySearch(names, target)

	if index != -1 {
		fmt.Printf("Our name '%s', found in position %d. \n", target, index)
	} else {
		fmt.Printf("Our name '%s' not found. \n", target)
	}
}
