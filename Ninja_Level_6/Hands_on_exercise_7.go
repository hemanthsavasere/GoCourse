package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	l := 10
	a := factorial
	fmt.Println(a(l))
}
