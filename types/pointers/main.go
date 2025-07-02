package main

import "fmt"

func increment(i *int) {
	*i++
}

func main() {
	i := 1
	fmt.Println(i)
	increment(&i)
	fmt.Println(i)

}


I'll write my code here