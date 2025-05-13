package main

import (
	"fmt"
)

func printSlice(title string, items []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(items); i++ {
		element := items[i]
		fmt.Println(element)
	}
}

func slice() {

	mySlice := []string{"screws", "nails", "hammer"}

	printSlice("First-batch", mySlice)

	mySlice = append(mySlice, "tree")
	mySlice = append(mySlice, "paper")
	printSlice("second-batch", mySlice)

	mySlice = mySlice[2:]
	printSlice("third-batch", mySlice)
}
