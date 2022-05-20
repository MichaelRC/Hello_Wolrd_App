package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	//Websites pages
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
