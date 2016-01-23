package main

import (
	"html/template"
	"log"
	"net/http"
)

func GeneratorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("templates/generator.html")

		if err != nil {
			log.Println("Failed to parse files")
		}
		t.Execute(w, nil)

	default:
		http.Error(w, "Method not allowed", 405)

	}
}
