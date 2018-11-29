package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s1 := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	//fmt.Printf("%T \n", s1)
	fmt.Println(s1)

}
