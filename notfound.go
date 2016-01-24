package main

import (
	"net/http"
	"html/template"
)

func customNotFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	t, _ := template.ParseFiles("templates/404.html")

	t.Execute(w, nil)
}
