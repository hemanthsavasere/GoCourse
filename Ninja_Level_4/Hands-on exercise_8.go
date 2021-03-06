package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	marks := make(map[string][]int)

	for i := 0; i < 3; i++ {
		marks["Hemanth"] = []int{99, 99, 99, 99}
	}

	fmt.Println(marks)

	for key, val := range m {
		fmt.Println("This is the record for", key)
		for i, v2 := range val {
			fmt.Println("\t", i, v2)
		}
	}
}
