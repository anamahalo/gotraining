package main

import (
    "log"
)

type User struct {
    Name string
    Email string
}

// struct types having ability to contain
// anonymous or embedded fields
// when embedding type into struct
// name of type acts as field name for what is then
// an embedded field

// declare new type and embed User type into it
type Admin struct {
    User
    Level string
}


func (u *User) Notify() error {
    log.Printf("User: Sending User Email To %s<%s>\n",
            u.Name,
            u.Email)

    return nil
}

type Notifier interface {
    Notify() error
}

func SendNotification(notify Notifier) error {
    return notify.Notify()
}

// declared new type called Admin and embedded User type
// within struct declaration
// NOTE: this is COMPOSITION not inheritance

func main() {
    admin := &Admin{
        User: User{
            Name: "Jane Doe",
            Email: "jane@example.com",
        },
        Level: "super",
    }
    // #1
    // SendNotification(admin)

    // #2
    // admin.User.Notify()

    // #3
    admin.Notify()
}

// Output
// User: Sending User Email To Jane Doe<jane@example.com>

// Admin type now implements interface thru promotion of methods
// from embedded User type

// If Admin type now contains fields and methods of User type,
// then where are they in relationship to the struct?
//   "When we embed a type, methods of that type become
//    methods of the outer type,
//    but when they are invoked,
//    receiver of the method is the INNER type,
//    NOT the outer one" -Effective Go (emphasis own)

// Since name of embedded type acts as a field name and
// embedded type exists as an inner type, can then make
// following method call:
// see #2
// admin.User.Notify()

// fields and methods are also promoted to OUTER type:
// see #3
// admin.Notify()



// Rules for inner type method set promotion:
// Given struct type S and type named T,
// promoted methods are included in method set of struct:

//  If S contains anonymous field T,
//  method sets of S and *S both include promoted methods
//  w/receiver T
//      when we embed a type, methods for embedded type
//      w/receivers that use VALUE
//      are promoted and avail
//      for calling by values AND pointers of outer type

//  Method set of *S also includes promoted methods
//  w/receiver *T
//      when we embed a type, methods for embedded type
//      w/receivers that user a POINTER
//      are promoted and avail
//      for calling by pointers of the outer type

//  S contains anonymous field *T, method sets of S
//  and *S both incl. promoted methods w/receiver T
//  or *T
//      when we embed pointer of the type, methods
//      for embedded type w/receivers that use both
//      values and pointers are promoted and avail
//      for calling by values and pointers of outer type

//  If S contains anonymous field T, method set of S
//  DOES NOT include promoted methods w/receiver *T
//      when we embed a type, methods for embedded type
//      w/receivers that use a pointer aren't promoted
//      for calling by values of the outer type

// See embedded_types2.go
