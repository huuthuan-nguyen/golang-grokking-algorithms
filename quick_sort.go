package main

import "fmt"

func quickSort(list []int) []int {
	var result []int
	if len(list) < 2 {
		return list
	} else {
		var pivot = list[0]
		var less, greater []int
		for i, v := range list {
			if i != 0 {
				if v <= pivot {
					less = append(less, v)
				} else {
					greater = append(greater, v)
				}
			}
		}
		result = append(result, less...)
		result = append(result, pivot)
		result = append(result, greater...)
		return result
	}
}

func main() {
	fmt.Println(quickSort([]int{3, 6, 1, 7, 9}))
}
