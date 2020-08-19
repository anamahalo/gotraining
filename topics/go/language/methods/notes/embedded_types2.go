package main

import (
    "log"
)

type User struct{
    Name string
    Email string
}

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

func (a *Admin) Notify() error {
    log.Printf("Admin: Sending Admin Email To %s<%s>\n",
            a.Name,
            a.Email)

    return nil
}

type Notifier interface {
    Notify() error
}

func SendNotification(notify Notifier) error {
    return notify.Notify()
}

func main() {
    admin := &Admin {
        User: User{
            Name: "Jane Doe",
            Email: "jane@example.com",
        },
        Level: "super",
    }

    // create value of Admin type and pass address
    // of value to SendNotification()
    // SendNotification(admin)

    // calling Notify() using outer type
    // Admin type's implementation
    // User type's implementation no longer promoted to outer type
    admin.Notify()
}

// THUS

// Would compiler throw error b/c we now had 2 implementations
// of the interface?
//  No, b/c when we use embedded type, unqualified type's name
//  acts as field name.
//      Has effect of fields and methods of embedded type
//      having unique name as inner type of the struct
//          So we can have inner and outer implementation
//          of the same interface w/each implementation being
//          unique and accessible

// If compiler accepted type declaration, how does
// the compiler determine which implementation to use for
// interface calls?
//  If outer type contains implementation that satisfies
//  the interface, it will be used.
//  Otherwise, thanks to method promotion,
//  any inner type that implements interface can be used
//  thorugh the outer type.
