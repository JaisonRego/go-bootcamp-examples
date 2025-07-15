package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	channel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// channel <- 3
	// fmt.Println(<-channel)
	wg.Add(2)
	go func(channel <-chan int, wg *sync.WaitGroup) {
		ch, channelIsOpen := <-channel
		if channelIsOpen {
			fmt.Println("Value in channel:", ch)
		} else {
			fmt.Println("Value is garbage:", ch)
		}
		wg.Done()
	}(channel, wg)

	go func(channel chan<- int, wg *sync.WaitGroup) {
		channel <- 1
		channel <- 2
		close(channel)
		wg.Done()
	}(channel, wg)
	wg.Wait()
}
