package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/printURL/", printURL)
	err := http.ListenAndServe(":2060", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func printURL(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, req.URL.Path)
}
