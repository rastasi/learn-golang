package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Golang")

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "hello\n")
	} else {
		fmt.Fprintf(w, "lofasz\n")
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
