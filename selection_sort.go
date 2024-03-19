package main

import "fmt"

// find the smallest value in array
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr { // loop and compare
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var result []int
	for range len(arr) {
		smallestIndex := findSmallest(arr)
		result = append(result, arr[smallestIndex])                 // append to result array
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...) // remove element from slice
	}
	return result
}

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
