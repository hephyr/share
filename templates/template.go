package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	// go.rice provides methods to add resources to a binary
	layoutsTemplateBox := rice.MustFindBox("layouts")
	widgetsTemplateBox := rice.MustFindBox("widgets")
	widgetsStrings := []string{}
	widgetsTemplateBox.Walk("", func(widgetPath string, info os.FileInfo, err error) error {
		if path.Ext(widgetPath) == ".html" {
			widgetString, err := widgetsTemplateBox.String(widgetPath)
			if err != nil {
				log.Fatal(err)
			}
			widgetsStrings = append(widgetsStrings, widgetString)
		}
		return nil
	})

	layoutsTemplateBox.Walk("", func(layoutPath string, info os.FileInfo, err error) error {
		if path.Ext(layoutPath) == ".html" {
			layoutString, err := layoutsTemplateBox.String(layoutPath)
			if err != nil {
				log.Fatal(err)
			}
			templateString := strings.Join(append(widgetsStrings, layoutString), "\n")
			filename := strings.TrimSuffix(filepath.Base(layoutPath), path.Ext(layoutPath))

			tmpl, err := template.New(filename).Parse(templateString)
			if err != nil {
				log.Fatal(err)
			}
			templates[filename] = tmpl
		}
		return nil
	})
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("%v The template %s does not exist. ????", templates, name)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, name, data)
}
