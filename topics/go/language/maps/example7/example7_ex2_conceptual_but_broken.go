// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps are reference types.
package main

import "fmt"

// set the constant here to see whether we can pass it
// into func main() instead
const multnum = 2

func main(multnum int) {

	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Pass the map to a function to perform some mutation.
	multiplier(scores, "anna", multnum)

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}

// double finds the score for a specific player and
// multiplies it by the set constant.
func multiplier(scores map[string]int, player string, multnum int) {
	scores[player] = scores[player] * multnum

// # command-line-arguments
// example7/example7_ex2_conceptual_but_broken.go:13:6: func main must have no arguments and no return values
}
