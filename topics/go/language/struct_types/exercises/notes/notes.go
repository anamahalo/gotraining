// Lacks inheritance and thus traditional polymorphism.
// OOP in Go
package main

import (
    "fmt"
)

type Animal struct {
	Name string
	mean bool
}

type AnimalSounder interface {
    MakeNoise()
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

// func (dog *Dog) MakeNoise() {
//   barkStrength := dog.BarkStrength
//
//   if dog.mean == true {
//       barkStrength = barkStrength * 5
//   }
//
//   for bark := 0; bark < barkStrength; bark ++ {
//       fmt.Printf("BARK ")
//   }
//
//   fmt.Println("")
// }
//
// func (cat *Cat) MakeNoise() {
//   meowStrength := cat.MeowStrength
//
//   if cat.Basics.mean == true {
//       meowStrength = meowStrength * 5
//   }
// 
//   for meow := 0; meow < meowStrength; meow ++ {
//       fmt.Printf("MEOW ")
//   }
//
//   fmt.Println("")
// }

func main () {
    myDog := &Dog {
        Animal {
            "Rover", // Name
            false, // mean
        },
        2, //BarkStrength
    }

    myCat := &Cat {
        Basics: Animal {
            Name: "Julius",
            mean: true,
        },
        MeowStrength: 3,
    }

    MakeSomeNoise(myDog)
    MakeSomeNoise(myCat)
}

func (animal *Animal) PerformNoise(strength int, sound string) {
    if animal.mean == true {
        strength = strength * 5
    }

    for voice := 0; voice < strength; voice ++ {
        fmt.Println(sound)
    }

    fmt.Println("")
}

func (dog *Dog) MakeNoise() {
    dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Cat) MakeNoise() {
    cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

func MakeSomeNoise(animalSounder AnimalSounder) {
    animalSounder.MakeNoise()
}
