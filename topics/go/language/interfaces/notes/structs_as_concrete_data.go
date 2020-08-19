package main

import "fmt"

// concrete data
type file struct {
    name string
}

// concrete data
type pipe struct {
    name string
}

func main () {
    var f file // value of type
    var p pipe // value of type

    fmt.Println(f, p)
}
