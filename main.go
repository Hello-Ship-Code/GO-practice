package main

import "fmt"

func sum(a int, b int) int {
	return (a + b)
}

func main() {

	fmt.Println("Hello world")

	// variables

	var abc int
	fmt.Println("The sum is", abc)

	a, b := 1, 5
	fmt.Println("A:", a, "B:", b)

	result := sum(a, b)
	fmt.Println("The sum of ", a, "+", b, "=", result)

	var (
		c = 10
		d = 20
	)

	fmt.Println("C:", c, "D:", d)

	var age int = 20
	fmt.Println("Your Age:", age)

	//string

	var name string = "peter"
	// fmt.Println("Enter your name: ")
	// fmt.Scan(&name)
	fmt.Println("Hello Mr.", name)
}
