package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go send(c)
	receive(c)
	fmt.Println("Exiting")
}

func send(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
