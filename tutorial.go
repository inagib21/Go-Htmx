package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Film represents a movie with its title and director information.
type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello there")

	// h1 is the handler for the main page.
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// Parse the HTML template.
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Prepare film data.
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}

		// Execute the template with the film data.
		tmpl.Execute(w, films)
	}

	// h2 is the handler for adding a new film.
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		// Retrieve film title and director from the form.
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// Parse the HTML template.
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Execute the "film-list-element" template with the new film data.
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	// Define routes and corresponding handlers.
	http.HandleFunc("/", h1)          // Main page route.
	http.HandleFunc("/add-film/", h2) // Add film route.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
