package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Println("Hello, my name is", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"Bob"}
	talker.Talk()
}
