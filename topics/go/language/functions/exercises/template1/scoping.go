package main

import (
    "fmt"
)

type user struct {}

func main() {
    user := &user{}
    fmt.Println(user)
    x := &user{}
    fmt.Println(user)
}

// Compilation Error:
// ./prog.go:12:12: user is not a type

// user is declared at the global scope
// but the main function gets its own scope
// when you create a var named "user" it's added to the inner scope
// so when you go to access the symbol "user" it consults
// the inner scope *first* and finds the var
// and the var is *not* a type so it errors
