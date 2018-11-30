package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Println(string(i))
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U\t%v\n", i, string(i))
		}
	}
}
