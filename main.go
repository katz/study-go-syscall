package main

import (
	"fmt"
	"os"
)

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
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
