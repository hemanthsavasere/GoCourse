package main

import (
	"fmt"
)

func sumVariadic(x ...int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}

func sum(x []int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}

func cleanup() {
	fmt.Println("Done with cleanup")
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := sumVariadic(arr...)
	l := sum(arr)
	fmt.Println(k, l, "are sum of the same array respectively")
	defer cleanup()
}
