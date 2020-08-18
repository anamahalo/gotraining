package main

import (
    "os"
    "fmt"
)

func main() {
    // first argument, i.e.: program name, is excluded
    argLength := len(os.Args[1:])
    // use fmt.Printf() to format string
    fmt.Printf("Arg length is %d\n", argLength)
    
    // iterate over CLI args
    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a)
    }
}
