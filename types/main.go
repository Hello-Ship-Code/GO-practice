package main

import "fmt"

type sample struct {
	word string
	a, b int
}

// var sampleOne struct {
// 	word string
// 	a, b int
// }

func main() {
	data := sample{}
	fmt.Println("default values:", data)

	data.word = "Peter"
	data.a, data.b = 1, 2
	fmt.Println("populated:", data)

	other := sample{"john", 3, 2}
	fmt.Println("other:", other)

	var demo sample
	demo.word = "test"
	demo.a, demo.b = 2, 3
	fmt.Println("Demo: ", demo)

	sampleTwo := struct {
		word string
		a, b int
	}{
		"hello",
		1, 2,
	}
	fmt.Println("Sample Two:", sampleTwo)

	// Arrays

	fmt.Println("Array output:")
	myArray := [...]int{1, 2, 3}
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Array: %d\n", myArray[i])
	}

}
