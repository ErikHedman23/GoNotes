package main

import (
	"fmt"
	"reposgo/gonotes/sums"
)

func main() {
	fmt.Println("Hello World")
	sums := sums.Sum(1, 2)

	fmt.Println(sums)
}