package main

import (
	"fmt"
)

func summ(array []int) int {
	sum1 := 0
	for _, val := range array {
		sum1 += val
	}
	return sum1
}
func evenSum(x []int, f func(arr []int) int) int {
	var arr []int
	for _, val := range x {
		if val%2 == 0 {
			arr = append(arr, val)
		}
	}
	return f(arr) // calling the function and returning the result of the function
}

func main() {
	l := []int{1, 2, 3, 4, 5}
	n := summ(l)
	k := evenSum(l, summ)
	fmt.Println(n, k)
}
