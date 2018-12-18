package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
		wg.Done()
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar ", i)
	}
}
