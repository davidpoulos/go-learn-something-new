package main

import "fmt"

func main() {

	fmt.Println("Begin Program")
	go func() {
		fmt.Println("I'm being run concurrently")
	}()

	fmt.Println("Exit Program")

	fmt.Scanln()
}
