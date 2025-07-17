package main

import (
	"fmt"
	"strings"
)

var storeEmail []string
var x int = 0

func checkStringLength(text string) {
	var isStringLongEnough = len(text)
	var isStringLegitEmail = strings.Contains(text, "@")
	var isStringLegitDomain = strings.Contains(text, ".com")
	if isStringLongEnough > 4 && isStringLegitEmail && isStringLegitDomain {
		fmt.Print("Your email is greater than 4 characters and legit accepted \n")
		storeEmail = append(storeEmail, text)
	} else {
		fmt.Printf("Your email %v is short and fake, kicked out", text)
	}
	x = x + 1

	fmt.Print(x, storeEmail, "\n")
}
