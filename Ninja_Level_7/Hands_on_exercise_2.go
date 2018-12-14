package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) ChangeMe(new1, new2 string) {
	(*p).first = new1
	(*p).last = new2
	(*p).age = 23 + 2
}

func main() {
	p1 := person{
		first: "Steve",
		last:  "Jobs",
		age:   23,
	}
	p1.ChangeMe("Hemanth", "Savasere")
	fmt.Println("First Name", p1.first)
	fmt.Println("Last Name", p1.last)
	fmt.Println("Age", p1.age)
}
