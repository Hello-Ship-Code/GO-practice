package main

import (
	"fmt"
)

func rangefunc() {
	slice := []string{"hello", "world", "!"}

	for i, element := range slice {
		fmt.Println(i, element) // prints each element in the array

		for _, ch := range element {
			fmt.Printf(" %q\n", ch) // print each character
		}
	}
}
