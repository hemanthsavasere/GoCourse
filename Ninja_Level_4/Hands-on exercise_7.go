package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"Jame", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "penny", "Hello, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	xxz := [][]string{
		{"Jame", "Bond", "Shaken, not stirred"},
		{"Miss", "penny", "Hello, James."},
	}

	fmt.Println(xxz)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
