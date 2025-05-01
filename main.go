package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}

func main() {

	fmt.Println("Hello world")

	// variables

	var abc int
	fmt.Println("Garbage value if no value assigned:", abc)

	a, b := 1, 5
	fmt.Println("A:", a, "B:", b)

	var (
		c = 10
		d = 20
	)

	fmt.Println("C:", c, "D:", d)

	var age int = 20
	fmt.Println("Your Age:", age)

	// functions

	result := sum(a, b)
	fmt.Println("The sum of ", a, "+", b, "=", result)

	ab, bc, _ := multiReturn()

	fmt.Println("AB:", ab, "BC:", bc)

	//string

	var name string = "peter"
	// fmt.Println("Enter your name: ")
	// fmt.Scan(&name)
	fmt.Println("Hello Mr.", name)

	// flow control

}
