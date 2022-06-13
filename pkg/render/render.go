package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

//A map of fuctions that can be used
//not built into the template 'language'
//but Go allows us to pass them in and use them this way.
var functions = template.FuncMap{}

//RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Erorr getting template cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//find anything that ends in .page.gohtml (* is wildcard)
	pages, err := filepath.Glob("./template/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		//extracts name of file
		name := filepath.Base(page)
		//print to console name of file grabbed
		fmt.Println("Page is currently", page)

		//creating a template set [ts]
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//check to see if there are any layouts that match this template
		matches, err := filepath.Glob("./template/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		//checks to see if # of matches is greater than 0
		//if TRUE parse the layout and match it to the template.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		//add template set to myCache map
		myCache[name] = ts
	}

	return myCache, nil
}
