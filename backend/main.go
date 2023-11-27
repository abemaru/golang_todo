package main

import (
	"fmt"
	"net/http"
)

func say_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", say_hello)
	http.ListenAndServe(":8888", nil)
}
