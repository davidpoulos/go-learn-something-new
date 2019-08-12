package main

import "fmt"

func main() {
	exampleOne()
	exampleTwo()
}

//exampleOne - Introducing Channels
func exampleOne() {

	// Initialize channel to send Strings to
	messaging := make(chan string)

	go func() {
		// Send value to a channel
		messaging <- "nacho cheese"
	}()

	// Receive any values that are sent to the channel
	// (Blocks until it receives a message)
	result := <-messaging

	fmt.Println(result)

	// NOTE!!!
	// By default sends and receives block until both the sender and receiver are ready.
}

//exampleTwo - Introducting Channel Buffers
func exampleTwo() {

	//By default channels are unbuffered, meaning that they will only accept sends (chan <-)
	//if there is a corresponding receive (<- chan) ready to receive the sent value.
	// Buffered channels accept a limited number of values without a corresponding
	// receiver for those values.
	cappedMessaging := make(chan string, 2)

	cappedMessaging <- "Ravioli"
	cappedMessaging <- "Spongebob"

	foodOne := <-cappedMessaging
	foodTwo := <-cappedMessaging

	fmt.Println(foodOne)
	fmt.Println(foodTwo)
}
