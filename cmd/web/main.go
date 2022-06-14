package main

import (
	"fmt"
	"github/MRC/firstgoweb/pkg/config"
	"github/MRC/firstgoweb/pkg/handlers"
	"github/MRC/firstgoweb/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

//main is the main application function
func main() {
	//Websites pages
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannont create template cache.")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
