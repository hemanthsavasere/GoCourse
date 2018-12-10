package main

import (
	"fmt"
)

type vehicle struct {
	doors  int
	colors []string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s1 := sedan{}
	s1.doors = 4
	s1.colors = make([]string, 1, 10)
	s1.colors[0] = "Red"
	s1.colors = append(s1.colors, "Green")
	fmt.Println(s1)
	for _, val := range s1.colors {
		fmt.Println(val)
	}
}
