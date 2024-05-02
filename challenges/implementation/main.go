package main

import (
	"challenges/imp/struct"
	"fmt"
	"log"
	"math"
)

type Squarey struct {
	Length float64
}

func (s Squarey) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// an interface defines a set of methods a type should implement:
type Shape interface {
	Area() float64
}

func main() {
	s, err := structs.NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())

	m := Squarey{20}
	fmt.Println(m.Area())

	c := Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{m, c}
	sa := sumAreas(shapes)

	fmt.Println(sa)
}
