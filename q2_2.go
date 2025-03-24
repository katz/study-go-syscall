package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("q2_2.txt")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	data := [][]string{
		{"Header1", "Header2", "Header3"},
		{"1.23", "Hello, World!", "123"},
		{"4.56", "Another String", "456"},
	}

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	writer.Flush()

	file.Close()
}
