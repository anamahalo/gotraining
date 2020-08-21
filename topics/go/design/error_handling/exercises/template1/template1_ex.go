// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create two error variables, one called ErrInvalidValue and the other
// called ErrAmountTooLarge. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the ErrInvalidValue. Check the value for greater than $1,000 and if it is,
// return the ErrAmountTooLarge. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

// Add imports.
import (
    "errors"
    "fmt"
)

var (
// Declare an error variable named ErrInvalidValue using the New
// function from the errors package.
    ErrInvalidValue = errors.New("Invalid Value")
// Declare an error variable named ErrAmountTooLarge using the New
// function from the errors package.
ErrAmountTooLarge = errors.New("Amount Too Large")
)

// Declare a function named checkAmount that accepts a value of
// type float64 and returns an error interface value.
// func checkAmount( /* parameter */ ) /* return arg */ {

	// Is the parameter equal to zero. If so then return
	// the error variable.

	// Is the parameter greater than 1000. If so then return
	// the other error variable.

	// Return nil for the error value.
// }
// See below for above function

func main() {

	// Call the checkAmount function and check the error. Then
	// use a switch/case to compare the error with each variable.
    if err := checkAmount(0); err != nil {
        switch err {
        case ErrInvalidValue:
            fmt.Println("Invalid value.")
            return
        case ErrAmountTooLarge:
            fmt.Println("Value is above threshold.")
            return
    // Add a default case. Return if there is an error.
        default:
            fmt.Println(err)
            return
        }
    }
	// Display everything is good.
    fmt.Println("Everything looks good.")
}

func checkAmount(f float64) error {
    if f == 0 {
        return ErrInvalidValue
    } else if f > 1000 {
        return ErrAmountTooLarge
    }

    return nil
}
