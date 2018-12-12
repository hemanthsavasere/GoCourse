package main

import (
	"fmt"
)

// Closure to linit scope of the function

func doubler() func() int {
	x := 1
	return func() int {
		x *= 2
		return x
	}
}

func main() {
	k := doubler()
	fmt.Println(doubler()())
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(k())

}
