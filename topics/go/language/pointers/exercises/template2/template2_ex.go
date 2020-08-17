// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import (
    "fmt"
)
// Declare a type named user.
type user struct {
    name string
    email string
}

// Create a function that changes the value of one of the user fields.
func funcName( /* add pointer parameter, add value parameter */ ) {

	// Use the pointer to change the value that the
	// pointer points to.
}

func main() {

	// Create a variable of type user and initialize each field.
    jane := user {
        name: "Jane",
        email: "jane@example.com"
    }

	// Display the value of the variable.
    fmt.Println("email:", jane.email)

	// Share the variable with the function you declared above.
    someFunc()

	// Display the value of the variable.
    fmt.Println("email:", jane.email)
}
