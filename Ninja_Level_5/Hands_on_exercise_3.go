package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	flavours []string
}

func main() {
	p1 := person{}
	p1.first = "Hemanth"
	p1.last = "Savasere"
	p1.flavours = make([]string, 2, 10)
	p1.flavours[0] = "Vanilla"
	p1.flavours[1] = "Chocolate"
	fmt.Println(p1)

	p2 := person{
		first:    "Lionel",
		last:     "Messi",
		flavours: []string{"Nutmeg", "Cross", "Markov"},
	}

	for _, val := range p2.flavours {
		fmt.Println(val)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.flavours {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
