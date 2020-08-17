// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import (
    "fmt"
)

// Add user type and provide comment.

// user represents a user in the system.
type user struct {
    name string
    email string
    age int
}

func main() {

	// Declare variable of type user and init using a struct literal.
    jane := user {
            name: "Jane",
            email: "jane@example.com",
            age: 100,
    }
	// Display the field values.
    fmt.Println("Name", jane.name)
    fmt.Println("Email", jane.email)
    fmt.Println("Age", jane.age)

	// Declare a variable using an anonymous struct.
    john := struct {
            name string
            email string
            age int
    }{
            name: "John",
            email: "john@example.com",
            age: 42
    }
	// Display the field values.
    fmt.Println("Name", john.name)
    fmt.Println("Email", john.email)
    fmt.Println("Age", john.age)
}
