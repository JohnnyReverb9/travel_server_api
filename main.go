package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

//go:embed assets/*
var assets embed.FS

var tpl = template.Must(template.ParseFS(templates, "templates/*.html"))

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", staticFileHandler()))
	mux.Handle("/assets/", http.StripPrefix("/assets/", renderAssets()))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	log.Println("Listening on http://localhost:" + port)

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Println("error starting server!")
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := tpl.ExecuteTemplate(w, tmpl+".html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderAssets() http.Handler {
	assetsFS, err := fs.Sub(assets, "assets")

	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(assetsFS))
}

func staticFileHandler() http.Handler {
	staticFS, err := fs.Sub(static, "static")

	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(staticFS))
}
