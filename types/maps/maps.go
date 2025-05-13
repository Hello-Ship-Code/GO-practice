package main

import (
	"fmt"
)

func maps() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 3
	shoppingList["biscuits"] += 2
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "biscuits")

	fmt.Println("delete: ", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	_, found := shoppingList["cereal"]
	if !found {
		fmt.Println("not found")
	}

	totalItems := 0

	for _, count := range shoppingList {
		totalItems += count
	}

	fmt.Println("total Items:", totalItems)

}
