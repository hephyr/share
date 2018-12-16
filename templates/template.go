package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templatesDir := "./templates/"
	// templatesDir := "./"
	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}
	widgets, err := filepath.Glob(templatesDir + "widgets/*.html")
	if err != nil {
		log.Fatal(err)
	}
	for _, layout := range layouts {
		files := append(widgets, layout)
		filename := strings.TrimSuffix(filepath.Base(layout), path.Ext(layout))
		templates[filename] = template.Must(template.ParseFiles(files...))
	}
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}
