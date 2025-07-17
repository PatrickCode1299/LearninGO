package main

import (
	"fmt"
)

func main() {
	var email string
	var i int = 11
	x := 0
	for x <= i {
		x++
		fmt.Print("What's is your email address?")
		fmt.Scan(&email)
		checkStringLength(email)
	}

}
