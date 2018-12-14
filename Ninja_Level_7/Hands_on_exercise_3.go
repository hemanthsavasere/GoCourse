package main

import (
	"fmt"
)

type stack struct {
	arr []int
}

func (s *stack) push(x int) {
	(*s).arr = append((*s).arr, x)
}

func (s *stack) pop() int {
	x := (*s).arr[len((*s).arr)-1]
	(*s).arr = (*s).arr[:len((*s).arr)-1]
	return x
}

func main() {
	s1 := stack{}
	s1.push(10)
	s1.push(20)
	s1.push(30)
	k := s1.pop()
	k = s1.pop()
	fmt.Println(k, s1)
}
