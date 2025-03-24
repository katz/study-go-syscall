package main

import (
	"fmt"
	"io"
	"os"
)

// see: https://mattn.kaoriya.net/software/lang/go/20140501172821.htm
type pascalCaseWriter struct {
	writer   io.Writer
	prevChar byte // 1つ前に書き込んだ文字
}

func (w *pascalCaseWriter) Write(p []byte) (int, error) {
	r := 0
	var b [1]byte
	for n, _ := range p {
		b[0] = p[n]
		switch w.prevChar {
		case ' ', '\t', '\r', '\n', 0:
			if 'a' <= b[0] && b[0] <= 'z' {
				b[0] -= 32
			}
		}

		nw, err := w.writer.Write(b[:])
		if err != nil {
			return r + nw, err
		}
		w.prevChar = b[0]
	}
	return r, nil
}

func NewPascalCaseWriter(w io.Writer) *pascalCaseWriter {
	return &pascalCaseWriter{w, 0}
}

func main() {
	w := NewPascalCaseWriter(os.Stdout)
	fmt.Fprintln(w, "hello world")
}
