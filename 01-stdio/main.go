package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter something to STDIN")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Printf("You Entered: %s", input)
}
