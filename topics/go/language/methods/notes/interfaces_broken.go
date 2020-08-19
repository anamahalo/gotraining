package main

import (
    "log"
)

type User struct {
    Name  string
    Email string
}

func (u *User) Notify() error {
    log.Printf("User: Sending User Email To %s<%s>\n",
            u.Name,
            u.Email)

    return nil
}

// convention to name interfaces with an -er suffix
// when interface contains only ONE method
// can specify as many methods; stdlib rarely 2+
type Notifier interface {
    Notify() error
}

func SendNotification(notify Notifier) error {
    return notify.Notify()
}

func main() {
    user := User {
            Name:  "Jane Doe",
            Email: "jane@example.com",
    }

    SendNotification(user)
}

// # command-line-arguments
// ./interfaces.go:37:21: cannot use user (type User) as type Notifier in argument to SendNotification:
// 	User does not implement Notifier (Notify method has pointer receiver)

// compiler does not consider our value to be of type
// that implements interface b/c:

//   method set of corresponding pointer type *T is
//   set of all methods w/receiver *T or T
//      if interface var we're using to call interface method
//      contains pointer, then methods w/receivers based
//      on both values and pointers will satisfy the interface
//        rule doesn't apply to our example b/c we're passing
//        a value to SendNotification()

//   method set of any other type T consists of all methods
//   w/receiver type T
//     if interface var we're using to call particular interface
//     method contains a value, then only methods w/receivers
//     based on VALUES will satisfy the interface
//       rule doesn't apply to our example b/c receiver for
//       Notify() accepts a pointer

// rule that applies to example (and thus compile error):
//    method set of corresponding type T DOES NOT consist
//    of any methods w/receiver type *T
//       Notify() using pointer for receiver and we're
//       using value to make interface method call
//       FIX: need to pass address of User value to
//       SendNotification function
//       see interfaces.go
