package main

import "fmt"

const (
	Online      = 0
	offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case offline:
			stats[offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Unhandled server Status")
		}
	}

	fmt.Println(stats[Online], "Servers are Online")
	fmt.Println(stats[offline], "Servers are offline")
	fmt.Println(stats[Maintenance], "Servers are Maintenance")
	fmt.Println(stats[Retired], "Servers are Retired")
}

func maping() {
	servers := []string{"darkstar", "Aiur", "omnicorn", "ChocoTaco", "base64"}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["Aiur"] = offline

	printServerStatus(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printServerStatus(serverStatus)
}
