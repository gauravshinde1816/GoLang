package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404  page nt found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Request not supported", http.StatusNotFound)
	}
	fmt.Fprint(w, "Hello world")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "POST successful")
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Println(name, email)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", homePageHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Print("Server i started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
