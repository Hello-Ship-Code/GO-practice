package main

import (
	"fmt"
	"strings"
)

func printOne(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat("*", n))
	}
}
