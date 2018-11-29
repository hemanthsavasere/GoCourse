package main

import (
	"fmt"
)

func main() {
	const (
		x int = 42
		y     = 55
	)
	fmt.Println(x, y)
	fmt.Printf("%T\t %T\t", x, y)
}
