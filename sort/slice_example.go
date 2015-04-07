/*
 * Block-style comment is for the package documentation.
 */
package main

import (
	"fmt"
)

func main() {
	input := make([]int, 12, 12)
	for i,_ := range input {
		input[i] = i
	}

	half := len(input)/2
	left := input[0: half]
	right := input[half:len(input)]
	
	test := input[3:5]

	fmt.Println("Unsorted input: ", input, " half: ", half)
	fmt.Println("Left: ", left, " right: ", right, " test: ", test)
}
