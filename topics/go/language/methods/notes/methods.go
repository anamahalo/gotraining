// Function declared w/a receiver
// Receiver is a value or pointer of a *named* or *struct* type
// All methods for a given type belong to the type's method set

// declare struct type and method for that type
// declare struct type User then declare method named Notify
// w/a receiver that accepts value of type User
type User struct (
    Name string
    Email string
)

func (u User) Notify() error

// to call Notify method we need a value or pointer of type User:

// value of type User can be used to call method
// w/a value receiver
jane := User("Jane", "jane@example.com")
jane.Notify()

// pointer of type User can also be used to call a method
// w/a value receiver
john := &User("John", "john@example.com")
john.Notify()

// ======================================
// when using pointer, Go adjusts and *dereferences* pointer so call
// can be made
// when receiver is NOT a pointer method is operating against
// a COPY of the receiver value

// change Notify method to use pointer for receiver:
func (u *User) Notify() error

// value of type User can be used to call method
// w/a POINTER receiver
jane := User("Jane", "jane@example.com")
jane.Notify()

// pointer of type User can be used to call a method
// w/a POINTER receiver
john := &User("John", "john@example.com")
john.Notify()
