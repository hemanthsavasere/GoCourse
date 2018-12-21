package main

// Stack program with Thread Safety and Also error handling

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
		mu.Unlock()
		wg.Done()
		return 0, fmt.Errorf("Error Elements cannot be popped as Stack is already emptied ")

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
	wg.Add(3)
	k, err := s1.pop()
	k, err = s1.pop()
	k, err = s1.pop()
	wg.Wait()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1, k)
	}

}
