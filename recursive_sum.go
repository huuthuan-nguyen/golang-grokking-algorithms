package main

import "fmt"

func sum(list ...int) int {
	if len(list) == 0 {
		return 0
	} else {
		return list[0] + sum(list[1:]...)
	}
}

func main() {
	var list []int = []int{1, 2, 3, 4}
	fmt.Println(sum(list...))
}
