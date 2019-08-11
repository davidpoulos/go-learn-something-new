package main

import "fmt"

type Creature interface {
	color() string
}

type Alien struct {
	colour string
}

//Color
func (a Alien) color() string {
	return a.colour
}

func createColor(c Creature) {
	fmt.Println(c.color())
}

func main() {
	yikes := Alien{"blue"}
	createColor(yikes)
}
