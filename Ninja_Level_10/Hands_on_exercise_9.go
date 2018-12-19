package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	v := send(c)
	for n := 0; n < 100; n++ {
		fmt.Println(n, <-v)
	}
	fmt.Println("Exiting")
}

func send(c chan<- int) <-chan int {
	v := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for k := 0; k < 10; k++ {
				v <- k
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	close(c)
	return v
}
