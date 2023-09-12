package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world in go! Port :8080")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is about page")
		log.Println("This is about page")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n%s", r.URL.Path[1:], html.EscapeString(r.URL.Path))
	log.Println(r.URL.Path)
}
