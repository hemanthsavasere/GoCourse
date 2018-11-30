package main

import (
	"fmt"
)

func main() {
	favsport := "Barcaleona"
	switch favsport {
	case "Barcaleona":
		fmt.Println("Lionel Messi")
	case "Real Madrid":
		fmt.Println("Christano Ronaldo")
	case "Paris Saint Germain":
		fmt.Println("Neymar Jr.")
	default:
		fmt.Println("your favourite footballer is ?")
	}
}
