package main

import "fmt"

type num int

func main() {
	var x num
	var y float32
	fmt.Printf("%T\n", x)
	y = float32(x)
	fmt.Println(y)
}
