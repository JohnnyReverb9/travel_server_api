package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"travel_server_api/structs"
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
	mux.HandleFunc("/users/{id}/visits", getUsersVisits)
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

// GET /users/{id}/visits : params: fromDate, toDate
func getUsersVisits(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.URL.Path, "/")

	if len(uri) < 4 {
		http.Error(w, "400 | Bad Request", http.StatusBadRequest)
		return
	}

	id := uri[2]
	var response []structs.VisitResponse
	var fromDate, toDate int64
	var err error

	queryParams := r.URL.Query()
	allowedParams := map[string]bool{
		"fromDate": true,
		"toDate":   true,
	}

	for param := range queryParams {
		if !allowedParams[param] {
			http.Error(w, "400 | Bad Request", http.StatusBadRequest)
			return
		}
	}

	if _, exists := users[id]; exists {
		if (r.URL.Query().Has("fromDate") && r.URL.Query().Get("fromDate") == "") || (r.URL.Query().Has("toDate") && r.URL.Query().Get("toDate") == "") {
			http.Error(w, "400 | Bad Request", http.StatusBadRequest)
			return
		}

		fromDateStr := r.URL.Query().Get("fromDate")
		toDateStr := r.URL.Query().Get("toDate")

		if fromDateStr != "" {
			fromDate, err = strconv.ParseInt(fromDateStr, 10, 64)
			if err != nil {
				http.Error(w, "400 | Bad Request", http.StatusBadRequest)
				return
			}
		}

		if toDateStr != "" {
			toDate, err = strconv.ParseInt(toDateStr, 10, 64)
			if err != nil {
				http.Error(w, "400 | Bad Request", http.StatusBadRequest)
				return
			}
		}

		for _, visit := range visits {
			if strconv.Itoa(int(visit.User)) == id {
				if (fromDateStr == "" || visit.Visited_at >= fromDate) && (toDateStr == "" || visit.Visited_at <= toDate) {
					location := locations[strconv.Itoa(int(visit.Location))]
					response = append(response, structs.VisitResponse{
						Mark:       visit.Mark,
						Visited_at: visit.Visited_at,
						Place:      location.Place,
					})
				}
			}
		}

		fullResponse := structs.VisitsResponse{Visits: response}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(fullResponse)

		if err != nil {
			log.Println(err)
		}
	} else {
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
