package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//A map of fuctions that can be used
//not built into the template 'language'
//but Go allows us to pass them in and use them this way.
var functions = template.FuncMap{}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		//closes the program and logs the error
		log.Println("Error 1", err)
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Println("Error 2", err)
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreareTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//find anything that ends in .page.gohtml (* is wildcard)
	pages, err := filepath.Glob("./template/*.page.gohtml")
	if err != nil {
		log.Println("Error CTC 1", err)
		return myCache, err
	}

	for _, page := range pages {
		//extracts name of file
		name := filepath.Base(page)

		//creating a template set [ts]
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Error CTC 2", err)
			return myCache, err
		}

		//check to see if there are any layouts that match this template
		matches, err := filepath.Glob("./template/*.layout.gohtml")
		if err != nil {
			log.Println("Error CTC 3", err)
			return myCache, err
		}

		//checks to see if # of matches is greater than 0
		//if TRUE parse the layout and match it to the template.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				log.Println("Error CTC 4", err)
				return myCache, err
			}
		}

		//add template set to myCache map
		myCache[name] = ts
	}

	return myCache, nil
}
