package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"time"
)

func main() {
	var num1 int = 1
	var num2 float64 = 4

	fmt.Println(num1 + int(num2))

	fmt.Println()
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random number is:", rand.Intn(5))

	fmt.Println(rand.Int(rand.Reader, big.NewInt(5)))

	presetTime := time.Now()
	fmt.Println(presetTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.December, 2, 2, 22, 22, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	a := new([]int)
	b := make([]int, 4)
	fmt.Println(a, b)

	b[0] = 123
	b[1] = 666
	b[2] = 333
	b[3] = 555

	b = append(b, 777)
	fmt.Println(b)
	fmt.Println(sort.IntsAreSorted(b))
	sort.Ints(b)
	fmt.Println(b)
	fmt.Println(sort.IntsAreSorted(b))
	b = append(b[:2], b[3:]...)
	fmt.Println(b)

	fmt.Println()
	c := User{"jaison", "jaison@co.in", 23, true}
	fmt.Println("The details are: ", c)
	fmt.Printf("The details are: %+v\n", c)
	c.getStatus()
	fmt.Println()

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}

func (u User) getStatus() {
	fmt.Println(u.status)
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
