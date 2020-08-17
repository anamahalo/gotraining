package main

import "fmt"

// Much larger value than int64
const myConst int64 = 9223372036854775808543522345

func main() {
    fmt.Println("Will NOT Compile")
}

// Compiler Error:
// ./notes_cannot_be_compiled.go:6:7: constant 9223372036854775808543522345 overflows int64
