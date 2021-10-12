package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

const location = "Remote"

var name string

func main() {
	name = "Johnny"
	from := `Haiti`
	var n int = 2

	var proverb = "Undefined"
	if p, err := proverbs.Nth(4); err == nil {
		proverb = p.Saying
	}

	fmt.Printf("Hello, fellow %s Gophers!\n", location)
	fmt.Printf("My name is %s and I'm from %s.\n", name, from)
	fmt.Printf("By the time %d o'clock EST comes around, we'll know how to code in Go!\\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
}
