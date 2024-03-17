package main

import "fmt"

// CheckBin searches for an integer in a sorted list uisng binary search.
// If returns the index of the element if found, otherwise returns -1.
func checkBin(list []int, i int) int {
	// Initialize the lower bound of the search range.
	low := 0

	// Initialize the upper bound of the search range.
	high := len(list) - 1

	// Perform binary serach until the search range is valid.
	for low <= high {
		// Calculate the middle index of the search range.
		mid := (low + high) / 2

		// If the middle element is equal to the target value, return its index.
		if list[mid] == i {
			return mid
		}

		// If the middle element is less than the target value, adjust the lower bound.
		if list[mid] < i {
			low = mid + 1
		} else {

			// If the middle element is greater than the target value, adjust the upper bound
			high = mid - 1
		}
	}

	// Return -1 if the target value is not found in the list
	return -1
}

func main() {
	// Define a sorted list of integers.
	mylist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Search for the value 17 in the list using binary search.
	// Print the index of the element if found, otherwise print -1.
	fmt.Println(checkBin(mylist, 17)) // Output: 16

	// Search for the value -1 in the list using binary search.
	// Print -1 since the value is not present in the list.
	fmt.Println(checkBin(mylist, -1)) // Output: -1
}
