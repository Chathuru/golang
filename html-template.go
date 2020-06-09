package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type htmlPage struct {
	Title string
	Body string
}

func index_hander(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Welcome to GO lang server!</h1>")
	p := htmlPage{Title: "Page 1", Body: "News 1"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func about_hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Ola, I am Chathura</h1>")
}

func main() {
	http.HandleFunc("/", index_hander)
	http.HandleFunc("/about", about_hander)
	http.ListenAndServe(":8080", nil)
}