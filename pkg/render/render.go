package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/OzodbekX/bookings/pkg/config"
	"github.com/OzodbekX/bookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {

}
func AddDefoultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmlp string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app != nil && app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmlp]
	if !ok {
		log.Fatal("could not get template")
	}

	buf := new(bytes.Buffer)
	td = AddDefoultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error in writing template in to browser")
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
