package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the array elements")
	values, _ := reader.ReadString('\n')

	var val []int
	for _, value := range []byte(values) {
		if value == 32 || value == 10 {
			continue
		}
		tval, _ := strconv.Atoi(string(value))
		val = append(val, tval)
	}
	sum, length, name := adder(val)
	fmt.Printf("The sum of the above array elements are %v having length %v and returning string %v\n", sum, length, name)
}

func adder(values []int) (int, int, string) {
	sum := 0
	for _, value := range values {
		sum = sum + value
	}
	length := len(values)
	name := "Just for fun"
	return sum, length, name
}
