package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// json 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// json 化してレスポンスに書き込む

	gw := gzip.NewWriter(w)

	mw := io.MultiWriter(gw, os.Stdout)

	enc := json.NewEncoder(mw)
	enc.SetIndent("", "  ")
	if err := enc.Encode(source); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	gw.Flush()

	gw.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
