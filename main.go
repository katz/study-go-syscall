package main

import (
	"fmt"
	"io"
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
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "MultiWriter example\n")
}
