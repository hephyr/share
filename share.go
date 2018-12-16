package main

import (
	"net/http"

	"share/serve"
)

func main() {
	http.ListenAndServe(":80", serve.FileServer(serve.Dir(".")))
	// http.ListenAndServe(":80", http.FileServer(http.Dir(".")))
}