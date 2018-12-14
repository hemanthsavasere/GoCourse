package main

import (
	"fmt"
	"sort"
)

type programmers struct {
	First string
	Last  string
	Age   int
}

type byAge []programmers

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type byLast []programmers

func (b byLast) Len() int {
	return len(b)
}

func (b byLast) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLast) Less(i, j int) bool {
	return b[i].Last < b[j].Last
}

func main() {
	p1 := programmers{
		First: "Hemanth",
		Last:  "Savasere",
		Age:   63,
	}

	p2 := programmers{
		First: "Ken",
		Last:  "Thompson",
		Age:   56,
	}

	p3 := programmers{
		First: "Dennis",
		Last:  "Ritchie",
		Age:   33,
	}

	ps := []programmers{p1, p2, p3}
	sort.Sort(byAge(ps))
	fmt.Println(ps)
	sort.Sort(byLast(ps))
	fmt.Print(ps)
}
