package main

import (
	"database/sql"
	"dnd/server"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

// TODO:
//	- un-stub the helper stuff
//	- Add
//		- resistances
//		- flaws
//		- bonds
//		- personality traits
//		- ideals
//
//	- host endpoint
//
//  - persistent storage
//
//	- Features:
//		- create images of characters based on their sheets
//		- add rolling functionality for each sheet

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "One-small-step-4-man"
	dbname   = "dnd"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// host static files
	// TODO: better directory structure?
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
