package main

import (
	"fmt"
)

type Products struct {
	price int
	name  string
}

func printStats(list [4]Products) {

	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}
	fmt.Println("Last item on the list:", list[totalItems-1])
	fmt.Println("Total Items:", totalItems)
	fmt.Println("Total Cost:", cost)

}

func arrays() {

	shoppingList := [4]Products{
		{1, "Apple"},
		{5, "shirt"},
		{10, "pant"},
	}

	printStats(shoppingList)

	shoppingList[3].price = 25
	shoppingList[3].name = "pig"

	printStats(shoppingList)

}
