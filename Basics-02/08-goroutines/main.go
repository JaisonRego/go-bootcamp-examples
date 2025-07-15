package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wait sync.WaitGroup
var mut sync.Mutex

var signals []string

func main() {
	// go greeter("Hello")
	// greeter("World")
	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://wrong.com",
		"https://google.com",
		"https://fb.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		wait.Add(1)
	}

	wait.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 1; i < 6; i++ {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(website string) {
	defer wait.Done()
	response, err := http.Get(website)
	if err != nil {
		fmt.Printf("Something went wrong for Website: %s giving status code: %d\n", website, 400)
	} else {
		mut.Lock()
		signals = append(signals, website)
		mut.Unlock()
		fmt.Printf("The Website: %s returned with Status Code: %d\n", website, response.StatusCode)
	}
}
