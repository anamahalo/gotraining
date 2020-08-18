package main

import (
    "os"
    "fmt"
)

func main() {
    // first arg is always program name
    programName := os.Args[0]
    fmt.Println(programName)
}
