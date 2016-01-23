package main

import (
	"fmt"
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
	
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Println("Failed to get post values")
		}

		rawText := r.PostFormValue("desc")
		fmt.Println(rawText)
		fmt.Fprintln(w, rawText)

	default:
		http.Error(w, "Method not allowed", 405)

	}
}
