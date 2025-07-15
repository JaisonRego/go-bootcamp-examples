package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name:")
	fullName, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v\n", strings.Title(strings.ToLower(fullName)))

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa center between 1 and 5:")
	ratingValue, _ := reader.ReadString('\n')
	ratings, _ := strconv.ParseFloat(strings.TrimSpace(ratingValue), 64)

	fmt.Printf("\n Thanks for the rating our dosa center with %v star rating\n\n Your rating was recorded in our syste, at %v\n\n", ratings, time.Now().Format(time.Stamp))

	switch ratings {
	case 5:
		fmt.Println("Bonus for team for a 5 star service")
	case 4, 3:
		fmt.Println("We are always improving")
	default:
		fmt.Println("Need Serious work on our side")

	}
}
