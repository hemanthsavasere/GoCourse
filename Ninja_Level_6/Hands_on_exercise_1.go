package main

import (
	"fmt"
)

func foo() int {
	return 10
}

func bar() (string, int) {
	return "Hello from  bar", 10
}

func main() {
	k := foo()
	l, n := bar()
	fmt.Println(k, l, n)
}
