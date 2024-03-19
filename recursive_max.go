package main

import "fmt"

func max(list ...int) int {
	if len(list) == 0 {
		return 0
	} else if len(list) > 1 {
		var maxNum int = max(list[1:]...)
		if maxNum >= list[0] {
			return maxNum
		}
	}

	return list[0]
}

func main() {
	var list []int = []int{1, 2, 7, 4, 5}
	fmt.Println(max(list...))
}
