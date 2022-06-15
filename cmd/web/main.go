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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//give our render access to app config
	//using '&' as a referance to a pointer
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
