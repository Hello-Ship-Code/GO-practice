package main

import (
	"fmt"
)

func main() {
	var n, choice int
	fmt.Println("Enter the size & your choice: ")
	fmt.Scanf("%d %d", &n, &choice)

	switch choice {
	case 1:
		printOne(n) // pass `n` to the function
	default:
		fmt.Println("Invalid choice")
	}
}
