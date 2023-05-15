package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	}
}

func main() {
	e := Engine{}
	log.Fatal(http.ListenAndServe(":9999", e))
}
