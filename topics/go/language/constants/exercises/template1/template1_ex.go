// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import (
    "fmt"
)

const (
// Declare a constant named server of kind string and assign a value.
    server = "192.168.29.1#53"
// Declare a constant named port of type integer and assign a value.
    port int16 = 5720
)

func main() {

	// Display the value of both server and port.
    fmt.Println(server)
    fmt.Println(port)
	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.

	// Display the value of the variable.
}
