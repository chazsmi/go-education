package main

import (
	"fmt"
	"math"
)

// START STRUCT OMIT
// A struct
type Circle struct {
	x float64
	y float64
	r float64
}

// END STRUCT OMIT

func main() {
	//We can intialise a struct in different ways
	// START IN OMIT
	var c Circle
	c := new(Circle)
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r)
	// END IN OMIT

	// START CALL METHOD OMIT
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.Area())
	// END CALL METHOD OMIT

	// START ET CALL OMIT
	type Android struct {
		Person Person
		Model  string
	}
	a := new(Android)
	a.Person.Talk()
	// END ET CALL OMIT

	// START IFA OMIT
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	// END IFA OMIT
}

// START METHODS OMIT
func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

// END METHODS OMIT

// START EMBEDDED TYPES OMIT
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// END EMBEDDED TYPES OMIT

// START IF OMIT
type Animal interface {
	Speak() string
}

// Any type that defines this method is said to satisfy the Animal interface.
// There is no implements keyword in Go.
// Whether or not a type satisfies an interface is determined automatically
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// END IF OMIT

// START ITT OMIT
func useFullFunction(param interface{}) interface{} {
	return param
}

// END ITT OMIT
