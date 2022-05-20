package main

import (
	"fmt"
	"github/MRC/firstgoweb/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	//Websites pages
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
