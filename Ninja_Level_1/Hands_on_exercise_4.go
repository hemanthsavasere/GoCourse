package main

import "fmt"

type num int

func main() {
	var x num
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
