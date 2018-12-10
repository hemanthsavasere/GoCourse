package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first    string
		last     string
		friends  map[string]int
		flavours []string
	}{
		first: "Hemanth",
		last:  "Savasere",
		friends: map[string]int{
			"Steve Jobs":     0,
			"Steve Wozniack": 1,
			"Bill Gates":     9,
		},
		flavours: []string{
			"Nutmeg",
			"Apple",
			"Banana"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for key, val := range p1.friends {
		fmt.Println(key, val)
	}

	for _, val := range p1.flavours {
		fmt.Println(val)
	}
}
