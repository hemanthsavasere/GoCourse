package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "is Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "is Buzz")
		} else {
			fmt.Println("No fizz or buzz")
		}
	}
}
