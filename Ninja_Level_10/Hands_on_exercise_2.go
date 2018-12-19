// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello")
// 	c := make(chan int)
// 	c <- 42
// 	fmt.Println(<-c)
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
