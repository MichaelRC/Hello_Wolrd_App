package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Referance Middle Code (not being used)
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

/*
NoSurf creates a new CSRFToken for each visitor as a cookie

	Includes peramiters for use using the nosurf package
*/
func NoSurf(next http.Handler) http.Handler {
	//Create a new token
	csrfHandler := nosurf.New(next)

	//cookie peramiters
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
