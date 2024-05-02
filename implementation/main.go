package main

import (
	"fmt"
	"gonotes/http"
	"gonotes/sums"
	"time"
)

type Budget struct {
	Campaign string
	Balance  float64
	Expires  time.Time
}

//This is like creating a constructor that initializes the struct like in other languages:

func NewBudget(campaign string, balance float64, expires time.Time) (*Budget, error) {
	if campaign == "" {
		return nil, fmt.Errorf("empty campaign")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger than 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		Campaign: campaign,
		Balance:  balance,
		Expires:  expires,
	}

	return &b, nil
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

// We are passing in a pointer receiver because we want to actually change the value of the field of the struct
func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func main() {
	fmt.Println("Hello World")
	sums := sums.Sum(1, 2)

	fmt.Println(sums)

	ctype, err := noteshttp.ContentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1.TimeLeft())

	fmt.Printf("%#v\n", b1)

	b1.Update(10.5)
	fmt.Println(b1.Balance)

	//if I want to initialize fields with certain values from a struct:

	b2 := Budget{Balance: 245.6, Campaign: "puppies"}

	fmt.Printf("%v\n", b2)

	//you can also define a variable as the original values of the struct

	var b3 Budget

	fmt.Printf("%#v\n", b3)

	fmt.Println(b3)

	expires := time.Now().Add(7 * 24 * time.Hour)

	b4, err := NewBudget("puppies", 32.2, expires)
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Printf("%#v\n", b4)
	}
}
