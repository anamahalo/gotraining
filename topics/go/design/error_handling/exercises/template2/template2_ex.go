// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a custom error type called appError that contains three fields, err error,
// message string and code int. Implement the error interface providing your own message
// using these three fields. Implement a second method named temporary that returns false
// when the value of the code field is 9. Write a function called checkFlag that accepts
// a bool value. If the value is false, return a pointer of your custom error type
// initialized as you like. If the value is true, return a default error. Write a main
// function to call the checkFlag function and check the error using the temporary
// interface.
package main

// Add imports.
import (
    "errors"
    "fmt"
)
// Declare a struct type named appError with three fields, err of type error,
// message of type string and code of type int.
type appError struct {
    err     error
    message string
    code    int
}
// Declare a method for the appError struct type that implements the
// error interface.
func (a *appError) Error() string {
    return fmt.Sprintf("App Error: %s\n Message: %s\nCode: %d\n", a.err, a.message, a.code)
}
// Declare a method for the appError struct type named Temporary that returns
// true when the value of the code field is not 9.
func (a *appError) Temporary() bool {
    if a.code != 9 {
        return true
    }
    return false
}
// just `return a.code !=9` is simpler but personally I think the above
// has less strain on one's mental model b/c you can just read it
// however, `return a.code !=9` might be more likely seen in Golang world


// Declare the temporary interface type with a method named Temporary that
// takes no parameters and returns a bool.
type temporary interface {
    Temporary() bool
}
// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
// func checkFlag( /* parameter */ ) /* return arg */ {

	// If the parameter is false return an appError.

	// Return a default error.
// }
// See below for function

func main() {

	// Call the checkFlag function to simulate an error of the
	// concrete type.
    if err := checkFlag(false); err != nil {
	// Check the concrete type and handle appropriately.
	    switch e := err.(type) {
	// Apply the case for the existence of the Temporary behavior.
    // Log the error and write a second message only if the
	// error is not temporary.
        case temporary:
            fmt.Println(err)
            if !e.Temporary() {
                fmt.Println("This is a terrible 2nd error message")
            }
    
	// Apply the default case and just log the error.
        default:
            fmt.Println("Default Error") // add this so we know what's output
            fmt.Println(err)
        }
    }
}

func checkFlag(t bool) error {
    if !t {
        return &appError{errors.New("This is the error message! Flag is False"), "Message", 42} 
    }

    // default error
    return errors.New("Flaggity")
}
