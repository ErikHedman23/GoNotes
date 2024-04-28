package main

import (
	"fmt"
	"gonotes/sums"
)

func main() {
	fmt.Println("Hello World")
	sums := sums.Sum(1, 2)

	fmt.Println(sums)
}
