package main

import (
	"dnd/server"
	"log"
	"net/http"
	"text/template"
)

func main() {

	// host static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// host main page on base url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// add handlers
	http.HandleFunc("/character", server.CharacterRouter)

	// start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
