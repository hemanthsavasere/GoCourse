package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func (p *person) speak() {
	fmt.Println("I am", (*p).First, (*p).Last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		First: "Hemanth",
		Last:  "Savasere",
		Age:   23,
	}
	saySomething(&p1)
	//saySomething(p1)
	fmt.Println(&p1)
}
