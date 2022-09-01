package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404  page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Request not supported", http.StatusNotFound)
	}
	fmt.Fprint(w, "<h1>Hello world</h1>")
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
	// check static folder ==> looks for index.html by default
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", homePageHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Print("Server is started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
