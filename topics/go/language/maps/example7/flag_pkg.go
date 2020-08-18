package main

import (
    "flag"
    "fmt"
)

func main() {
    // var declaration
    var uname string
    var pass string

    // flags declaration using flag pkg
    flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
    flag.StringVar(&pass, "p", "password", "Specify pass. Default is password")

    // after declaring flags we need to call it
    flag.Parse()

    // check if CLI params match
    if uname == "root" && pass == "password" {
        fmt.Println("Logging in 🥂")
    } else {
        fmt.Println("Invalid credentials!")
    }
}

// !!problem!!
// presently anything can be passed in and it will display logging in
//   example:
//   ./flag_pkg foobar
//   returns "Logging in"
// need some kind of acceptable flags only validation check
