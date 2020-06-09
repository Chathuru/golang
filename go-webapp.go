package main

import (
	"fmt"
	"net/http"
)

func index_hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to GO lang server!</h1>")
}

func about_hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Ola, I am Chathura</h1>")
}

func main() {
	http.HandleFunc("/", index_hander)
	http.HandleFunc("/about", about_hander)
	http.ListenAndServe(":8000", nil)
}