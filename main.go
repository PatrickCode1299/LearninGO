package main

import (
	"fmt"
	"strings"
)

func main() {

	printMyName()
}
func printMyName() {

	var multiNames [50]string
	multiNames[0] = "John Tom"
	multiNames[1] = "Patrick Daniels"
	multiNames[2] = "Thomas Samuel"
	var singleNames []string

	for _, n := range multiNames {
		var names = strings.Fields(n)
		var singleNames = append(singleNames, names[0])
		fmt.Print(singleNames, "\n")

	}

}
