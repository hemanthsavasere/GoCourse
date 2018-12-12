package main

import (
	"fmt"
)

func main() {
	k, l := 10, 20
	k, l = func(x, y int) (int, int) {
		x, y = y, x
		return x, y
	}(10, 20)
	fmt.Println(k, l)
}
