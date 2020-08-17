// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

// Add imports.
import (
    "fmt"
)

func main() {

	// Declare an integer variable with the value of 20.
    intVar := 20

    // Display the address of and value of the variable.
    fmt.Println("Address Of:", &intVar, "Value Of:", intVar)

	// Declare a pointer variable of type int. Assign the
    // address of the integer variable above.
    pointer := &intVar

	// Display the address of, value of and the value the pointer
	// points to.
    fmt.Println("AddressOf:", &pointer, "Value Of:", pointer, "Points To:", *pointer)
}
