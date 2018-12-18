package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var counter int64
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			mu.Lock()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumCPU())
	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("count", counter)
}
