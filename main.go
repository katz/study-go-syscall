package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// example.comにアクセスして、そのレスポンスをstdoutに書き出す
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")

	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Header)

	io.Copy(os.Stdout, res.Body)
}
