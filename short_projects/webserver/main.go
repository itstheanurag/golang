package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println(" Starting the webser at the Port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("could not start the server")
	}
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)

		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)

		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post Request Successful")

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "name() err: %v", name)
	fmt.Fprintf(w,"email() err: %v", email)
	fmt.Fprintf(w, "password() err: %v", password)

}