// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how maps are reference types.
package main

import "fmt"

// need to utilize flag or os pkg if want to take in CLI commands
// to set the constant
// revisit this later
func main() {

	// Initialize a map with values.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

    // set a constant for whatever num we want to multiply by
    const multnum = 2

	// Pass the map to a function to perform some mutation.
	multiplier(scores, "anna", multnum)

	// See the change is visible in our map.
	fmt.Println("Score:", scores["anna"])
}

// double finds the score for a specific player and
// multiplies it by the set constant.
func multiplier(scores map[string]int, player string, multnum int) {
	scores[player] = scores[player] * multnum
}
