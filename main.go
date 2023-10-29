package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Contact struct {
	Email     string
	FirstName string
	LastName  string
}

func main() {
	// hardcoded contacts data
	contacts := map[string][]Contact{
		"Contacts": {
			{
				Email:     "rhuel@test.com",
				FirstName: "Rhuel",
				LastName:  "Garza",
			},
			{
				Email:     "garza@test.org",
				FirstName: "Bob",
				LastName:  "Bobbington",
			},
		},
	}

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// set headers
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "GET" {

			// get search string from request
			tmpl := template.Must(template.ParseFiles("./index.html"))

			search := strings.ToLower(r.URL.Query().Get("search"))
			log.Println("search query: ", search)

			// create a map to hold the filtered contacts
			filtered := make(map[string][]Contact)

			// filter out contacts for first matching email, fname, or lname
			for i := range contacts["Contacts"] {
				em := strings.ToLower(contacts["Contacts"][i].Email)
				fn := strings.ToLower(contacts["Contacts"][i].FirstName)
				ln := strings.ToLower(contacts["Contacts"][i].LastName)
				if strings.Contains(em, search) || strings.Contains(fn, search) || strings.Contains(ln, search) {
					filtered["Contacts"] = append(filtered["Contacts"], contacts["Contacts"][i])
				}
			}

			log.Println("filtered matches: ", filtered)

			if len(filtered) == 0 {
				log.Println("no found match for ", search)
				tmpl.Execute(w, nil)
			} else {
				log.Println("found match for ", search)
				tmpl.Execute(w, filtered)
			}
		}
	})

	log.Println("App started on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
