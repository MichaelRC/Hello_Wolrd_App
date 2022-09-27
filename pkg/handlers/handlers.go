package handlers

import (
	"github/MRC/firstgoweb/pkg/config"
	"github/MRC/firstgoweb/pkg/render"
	"net/http"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //security token 'Cross Site Request Forgery Token' for forms
	Flash     string //'flash message' posted to used (such as "success!")
	Warning   string
	Error     string
}

// Repo is the reposityory used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new reporistory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &TemplateData{})
}

// About is the about page handler.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	//send the data to template
	render.RenderTemplate(w, "about.page.gohtml", &TemplateData{
		StringMap: stringMap,
	})
}
