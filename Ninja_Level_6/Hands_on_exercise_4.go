package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi I am", p.first, p.last)
	fmt.Println("Nice to meet you buddy. I am ", p.age, "years old.")
}

func main() {
	p1 := person{}
	p1.first, p1.last, p1.age = "Hemanth", "Savasere", 23
	p1.speak()
}
