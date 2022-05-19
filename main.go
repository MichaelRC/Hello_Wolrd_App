package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

//About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {

}

//main is the main application function
func main() {
	//Websites pages
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
