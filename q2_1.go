package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("q2_1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "Float : %.2f\n", 1.23)
	fmt.Fprintf(file, "String: %s\n", "Hello, World!")
	fmt.Fprintf(file, "Int   : %d\n", 123)

	file.Close()
}
