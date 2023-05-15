package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	header := req.Header
	for k, v := range header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func indexHandler(writer http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(writer, "URL.Path = %q\n", req.URL.Path)
}
