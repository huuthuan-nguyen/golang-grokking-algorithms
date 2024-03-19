package main

import "fmt"

func count(list ...int) int {
	if len(list) == 0 {
		return 0
	} else {
		return 1 + count(list[1:]...)
	}
}

func main() {
	var list []int = []int{1, 2, 3, 4}
	fmt.Println(count(list...))
}
