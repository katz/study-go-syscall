package main

import (
	"io"
	"os"
)

func main() {
	// 指定したファイルを開き、内容をstdoutに出力する
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
