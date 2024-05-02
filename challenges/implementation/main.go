package main

import (
	"challenges/imp/struct"
	"fmt"
	"log"
)

func main() {
	s, err := structs.NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
