package main

import (
	"fmt"
)

func main() {
	var lifebooster float64 = 99.2
	lifeboosterRef := &lifebooster

	lifebooster = lifebooster * 2.2

	fmt.Println(lifebooster)
	fmt.Println(*lifeboosterRef)

}
