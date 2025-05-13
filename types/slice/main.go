package main

import (
	"fmt"
)

func printMySlice(title string, routes []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(routes); i++ {
		element := routes[i]
		fmt.Println(element)
	}

}

func main() {

	route := []string{"hotel", "park", "airport"}

	printMySlice("Route 1", route)
	route = append(route, "Home")
	printMySlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printMySlice("Remaining locations", route)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Second-program")
	slice()

}
