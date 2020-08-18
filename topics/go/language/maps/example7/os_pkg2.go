package main

import (
    "os"
    "fmt"
)

func main() {
    argLength := len(os.Args[1:])
        // use fmt.Printf() to format string
    fmt.Printf("Arg length is %d\n", argLength)
}
