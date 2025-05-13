package main

import (
	"fmt"
)

type Room struct {
	name    string
	cleaned bool
}

func checkClean(rooms []Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {

	rooms := []Room{
		{name: "office"},
		{name: "warehouse"},
		{name: "washroom"},
		{name: "ops"},
		{name: "basement"},
	}

	checkClean(rooms)

	fmt.Println("Cleaning rooms...")

	rooms[2].cleaned = true
	rooms[3].cleaned = true
	rooms[4].cleaned = true

	checkClean((rooms))

	// Arrays

	// arrays()

	//range

	// rangefunc()
}
