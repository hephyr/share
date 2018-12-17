package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"share/serve"
)

var dir string

func init() {
	flag.StringVar(&dir, "d", "", "set share directory")
	flag.Parse()
}

func main() {
	if dir == "" {
		dir = flag.Arg(0)
	}
	if dir != "" {
		if !filepath.IsAbs(dir) {
			currentDir, _ := os.Getwd()
			dir = filepath.Join(currentDir, dir)
		}
	}
	http.Handle("/", http.RedirectHandler("/files/", http.StatusMovedPermanently))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/files/", http.StripPrefix("/files/", serve.FileServer(serve.Dir(dir))))
	http.ListenAndServe(":80", nil)
}
