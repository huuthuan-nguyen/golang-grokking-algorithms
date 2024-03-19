package main

import (
	"fmt"
)

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2 // this is to prevent overflow

		// check if the item is present at mid
		if list[mid] == item {
			return mid
		}

		// if item is greater, ignore left half
		if list[mid] < item {
			low = mid + 1
		} else {
			// if item is smaller, ignore right half
			high = mid - 1
		}
	}

	// if we reach here, then the element was not present
	return -1
}

func main() {
	list := []int{2, 3, 4, 10, 40}
	item := 10
	// do binary search
	result := binarySearch(list, item)
	if result != -1 {
		fmt.Printf("Item found at index %d\n", result)
	} else {
		fmt.Println("Item not found in the array")
	}
}
