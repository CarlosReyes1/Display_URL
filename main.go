package main

import (
	"io"
	"net/http"
)

type Shoot int

func (h Shoot ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {
	var gun Shoot

	mux := http.NewServeMux()
	mux.Handle("/", gun)

	http.ListenAndServe(":8080", mux)
}
