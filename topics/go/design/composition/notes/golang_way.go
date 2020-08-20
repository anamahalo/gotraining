// Focusing on common behavior
// Reference: https://www.ardanlabs.com/blog/2016/10/reducing-type-hierarchies.html
package main

import "fmt"

// Speaker provide a common behavior for all concrete types
// to follow if they want to be part of this group.
// This is a contract for these concrete types to follow.
type Speaker interface {
    Speak()
}

// Dog contains everything a Dog needs
type Dog struct {
    Name       string
    IsMammal   bool
    PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() {
    fmt.Println("Woof!",
        "My name is", d.Name,
        ", it is", d.IsMammal,
        "I am a mammal w/a pack factor of", d.PackFactor)
}

// Cat contains everything a Cat needs.
type Cat struct {
    Name        string
    IsMammal    bool
    ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak
func (c Cat) Speak() {
    fmt.Println("Meow!",
        "My name is", c.Name,
        ", it is", c.IsMammal,
        "I am a mammal w/a climb factor of", c.ClimbFactor)
}

// ============================================================
// Note how from oo_way.go Animal type is removed and copies
// common fields directly into Dog and Cat.
// Why?
//  Animal type as providing abstraction layer of reusable state
//  Program NEVER needed to creat or solely use value
//  of type Animal
//  Implementation of Speak method for Animal type was
//  a generalization
//  Speak method for the Animal type was never going to
//  be called

// Guidelines to declaring types:
//  Declare types that represent something new or unique
//  Validate that a value of any type is created or used on its own
//  Embed types to REUSE EXISTING BEHAVIORS you NEED to satisfy
//  Question types that are an alias or abstraction for an existing type
//  Question types whose sole purpose is to share common state
// ============================================================

func main() {
    speakers := []Speaker{
        // Create a Dog by initializing its Animal parts
        // and then its specific Dog attributes.
        Dog {
            Name:     "Fido",
            IsMammal: true,
            PackFactor: 5,
        },

        // Create a Cat by initializing its Animal parts
        // and then its specific Cat attributes
        Cat {
            Name:     "Milo",
            IsMammal: true,
            ClimbFactor: 4,
        },
    }

    // Have the Animals speak
    for _, spkr := range speakers {
        spkr.Speak()
    }
}

// THUS
// Didn't need base type or type hierarchies to group
// concrete values together
// Interface allowed us to create a slice of different
// concrete type values and work WITH these values
// THROUGH their COMMON BEHAVIOR
// Removed any type pollution by NOT declaring a type
// that was never solely used by the program
