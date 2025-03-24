package main

import (
	"io"
	"net"
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
	io.Copy(os.Stdout, conn)
}
