package main

import "fmt"

func main() {
	// Initialize
	pirate := make(map[string]string)

	pirate["davey_jones"] = "locker"

	fmt.Println(pirate)

	ships := map[int]string{1: "black pearl"}

	fmt.Println(ships)
	fmt.Println(ships[1])

}
