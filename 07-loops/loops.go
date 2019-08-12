package main

import "fmt"

func main() {
	characters := make([]string, 10)

	characters = append(characters, "Harry Potter", "Hermione Granger", "Ronald Weasley")

	i := 0

	// Standard Iteration (While loop)
	for i < len(characters) {
		if len(characters[i]) == 0 {
			i++
			continue
		}

		fmt.Println("Character => " + characters[i])
		i++
	}

	// Infinite Loops
	for {
		break // Break so I don't crash
	}

	// For loop
	for i := 0; i < len(characters); i++ {
		fmt.Printf("Count => %d\n", i)
	}

	// Using RANGE
	for _, person := range characters {
		fmt.Println("Character => " + person)
	}

	cities := map[int]string{
		0: "Austin",
		1: "Dallas",
	}

	for key, city := range cities {
		fmt.Printf("Key => %d City => %s\n", key, city)
	}

	fmt.Println(len(characters))
}
