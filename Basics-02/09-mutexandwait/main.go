package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition!!")
	score := []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("First G")
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("Second G")
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.Lock()
		fmt.Println("Third G")
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		mut.RLock()
		fmt.Println("Read G")
		fmt.Println(score)
		mut.RUnlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
