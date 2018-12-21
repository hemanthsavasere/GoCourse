package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

type stack struct {
	arr []int
}

func (s *stack) push(num int) {
	mu.Lock()
	(*s).arr = append((*s).arr, num)
	mu.Unlock()
	wg.Done()
}

func (s *stack) pop() (int, error) {
	mu.Lock()
	if len((*s).arr) == 0 {
		return 0, fmt.Errorf("Error Stack cannot be emptied ")
	}
	popped := (*s).arr[0]
	(*s).arr = (*s).arr[1:]

	mu.Unlock()
	wg.Done()

	return popped, nil

}

func main() {
	s1 := stack{}
	wg.Add(2)
	go s1.push(10)
	go s1.push(20)
	wg.Wait()
	wg.Add(1)
	k, err := s1.pop()
	wg.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1, k)

}
