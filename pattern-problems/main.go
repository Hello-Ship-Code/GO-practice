package main

import (
	"fmt"
	"strings"
)

func _(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func pat2(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

func pat3(n int) {

	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", n-i+1))
	}

}

func pat4(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

}

func pat5(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func pat6(n int) {
	for i := 0; i < n; i++ {

	}
}

func pat7(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j < n-i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	n := 5

	pat7(n)
}
