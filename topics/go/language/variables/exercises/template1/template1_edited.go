// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var age int
	var name string
	var boolean bool

	// Display the value of those variables.
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(boolean)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	today := 17
	dayOfWeek := "What happened to Monday?"
	camelDay := false

	// Display the value of those variables.
	fmt.Println(today)
	fmt.Println(dayOfWeek)
	fmt.Println(camelDay)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	// Type, value
	fmt.Printf("%T [%v]\n", pi, pi)
}
