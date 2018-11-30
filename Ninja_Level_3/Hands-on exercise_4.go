package main

import (
	"fmt"
)

func main() {
	i := 1995
	for {
		if i > 2018 {
			break
		}
		fmt.Println(i)
		i++
	}
}
