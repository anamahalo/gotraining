package main

import "fmt"

// declared interfaces
type Reader interface {
    Read()
}

type Writer interface {
    Write()
}

type ReadWriter interface {
    Reader
    Writer
}

// .: can implement concrete type implementing
//    all three interfaces
type system struct {
    Host string
}

func (*system) Read()  { /* ... */ }
func (*system) Write() { /* ... */ }

// interface values still valueless
// i.e.: rw and r are not real here
// .: code can't assign the interface values to e/o
//    can only assign concrete data stored INSIDE
//    the interface values
func main() {
    var rw ReadWriter = &system{"127.0.0.1"}
    var r Reader = rw
    fmt.Println(rw, r)
}

// OUTPUT
// &{127.0.0.1} &{127.0.0.1}
// .: compiler valides whether concrete data inside of
//    one interface also satisfies the other
// REMEMBER to focus on RELATIONSHIPS interfaces have
// WITH CONCRETE DATA rather than implementation details
