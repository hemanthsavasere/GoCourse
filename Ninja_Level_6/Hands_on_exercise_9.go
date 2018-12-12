package main

import "fmt"

func factorialSum(x int) func() int {
	var arr []int
	fact := 1
	for i := 2; i <= x; i++ {
		fact *= i
		arr = append(arr, fact)
	}
	return func() int {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		return sum
	}
}

func main() {
	num := 10
	k := factorialSum(num)
	fmt.Println(k, k())
}
