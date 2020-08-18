package main

import (
    "flag"
    "fmt"
)

func main() {
    // var declaration
    var name string

    flag.StringVar(&name, "n", "admin", "Specify name. Default is admin")

    // after declaring flags we need to call it
    flag.Usage = func() {
        fmt.Println("Usage of our Program:")
        fmt.Println("./flag_pkg2 -n username\n")
        // flag.PrintDefaults() // prints default usage
    }

    flag.Parse()
}

// !!problem!!
// same as flag_pkg wherein anything can be passed and it seems
// to run, but didn't test this one for a success case
// prob needs some kind of error handling here
