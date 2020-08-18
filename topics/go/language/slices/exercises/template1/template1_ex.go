// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
    var numbers []int
	// Append numbers to the slice.
    for i := 0; i < 16; i++ {
        numbers = append(numbers, i*16)
    }
	// Display each value in the slice.
    for _, number := range numbers {
        fmt.Println(number)
    }

    fmt.Println("============")
    // Declare a slice of strings and populate the slice with names.
    // var strings []string
    names := []string{"Lincoln", "Roosevelt", "Washington", "Jackson"}
	// Display each index position and slice value.
    for i, name := range names {
        fmt.Printf("Index: %d Name: %s\n", i, name) // pkg notation here
    }
    fmt.Println("==============")
	// Take a slice of index 1 and 2 of the slice of strings.
    slice := names[1:3] // index : length // beg. and end
	// Display each index position and slice values for the new slice.
    for i, name := range slice {
        fmt.Printf("Index: %d Name: %s\n", i, name)
    }
}
