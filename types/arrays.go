package main

import "fmt"

func arrays() {
	myArray := [...]int{1, 2, 3}
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("%d\n", myArray[i])
	}
}
