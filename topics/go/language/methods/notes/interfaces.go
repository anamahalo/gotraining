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

type Notifier interface {
    Notify() error
}

func SendNotification(notify Notifier) error {
    return notify.Notify()
}

func main () {
    user := &User{
        Name: "Jane Doe",
        Email: "jane@example.com",
    }

    SendNotification(user)
}
