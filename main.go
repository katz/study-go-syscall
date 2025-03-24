package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// stdinからデータ読み込むコードを書く

	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}
		fmt.Printf("size=%d, buffer=%s\n", size, string(buffer))
	}
}
