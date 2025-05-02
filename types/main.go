// main.go
package main

import (
	"fmt"
)

type sample struct {
	word string
	a, b int
}

// anonymous Struct
var sampleOne struct {
	word string
	a, b int
}

func main() {
	data := sample{}
	fmt.Println("default values:", data)

	// declaring type 1
	data.word = "Peter"
	data.a, data.b = 1, 2
	fmt.Println("populated:", data)

	// declaring type 2
	other := sample{"john", 3, 2}
	fmt.Println("other:", other)

	// declaring type 3
	var demo sample
	demo.word = "test"
	demo.a, demo.b = 2, 3
	fmt.Println("Demo: ", demo)

	// anonymous struct
	sampleTwo := struct {
		word string
		a, b int
	}{
		"hello",
		1, 2,
	}
	fmt.Println("Sample Two:", sampleTwo)

	// Calling the array function
	fmt.Println("array from:", array())
}
