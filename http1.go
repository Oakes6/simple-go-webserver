package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func mirror(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", mirror)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
