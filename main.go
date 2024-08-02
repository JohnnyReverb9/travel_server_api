package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"travel_server_api/test"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

//go:embed assets/*
var assets embed.FS

var tpl = template.Must(template.ParseFS(templates, "templates/*.html"))

// TEST DATA
var users = test.Users
var locations = test.Locations
var visits = test.Visits

// END TEST DATA

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", staticFileHandler()))
	mux.Handle("/assets/", http.StripPrefix("/assets/", renderAssets()))
	mux.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})
	mux.HandleFunc("/users/", getEntityHandler)
	mux.HandleFunc("/visits/", getEntityHandler)
	mux.HandleFunc("/locations/", getEntityHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 | Not found", http.StatusNotFound)
		return
	})

	log.Println("Listening on http://localhost:" + port + "/main")

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Println("error starting server!")
		log.Fatal(err)
	}
}

// GET /{entity}/{id} -- entity: users, locations, visits
func getEntityHandler(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.URL.Path, "/")

	if len(uri) != 3 {
		http.Error(w, "404 | Not found", http.StatusNotFound)
		return
	}

	entity := uri[1]
	id := uri[2]

	switch entity {
	case "users":
		if user, exists := users[id]; exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Println(err)
				return
			}
			return
		} else {
			http.Error(w, "404 | Not found", http.StatusNotFound)
		}
	case "locations":
		if location, exists := locations[id]; exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(location)
			if err != nil {
				log.Println(err)
				return
			}
			return
		} else {
			http.Error(w, "404 | Not found", http.StatusNotFound)
		}
	case "visits":
		if visit, exists := visits[id]; exists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(visit)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			http.Error(w, "404 | Not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "404 | Not found", http.StatusNotFound)
	}
}

// HTML template renderer
func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := tpl.ExecuteTemplate(w, tmpl+".html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// assets renderer
func renderAssets() http.Handler {
	assetsFS, err := fs.Sub(assets, "assets")

	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(assetsFS))
}

// static files renderer
func staticFileHandler() http.Handler {
	staticFS, err := fs.Sub(static, "static")

	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(http.FS(staticFS))
}
