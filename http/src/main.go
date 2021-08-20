package main

import (
	"fmt"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":8000", nil)
}
