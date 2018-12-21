package main

import (
	"fmt"
) 

type stack struct {
	info string
}

type (s stack) Error() string{
	return fmt.Sprintf("here is the error as given %v", s.info)

}

func main() {
	c1 := stack{
		info :" I am the best"
	}
}