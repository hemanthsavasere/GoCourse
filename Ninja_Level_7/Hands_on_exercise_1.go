package main

import (
	"fmt"
)

func main() {
	a := 10
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", &a)
	fmt.Printf("%T \n", &a)
	fmt.Printf("%v \n", arr)
	fmt.Printf("%T \n", arr)
	fmt.Printf("%v \n", &arr)
	fmt.Printf("%T \n", &arr)
}
