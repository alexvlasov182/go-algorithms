package main

import "fmt"

func checkBin(list []int, i int) int {
	// min border of my numbers
	low := 0

	// Last number in the array because the slice start from 0...n
	high := len(list) - 1

	// If the guess less than or equal of our nunbers
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return mid
		}
		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	mylist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(checkBin(mylist, 17)) //
	fmt.Println(checkBin(mylist, -1)) // nil
}
