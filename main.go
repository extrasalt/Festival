package main

import (
"fmt"
"net/http"
)

func main(){
	http.HandleFunc("/new",NewPageHandler)
	http.ListenAndServe(":3001", nil)

}

func NewPageHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
  case "POST":
    fmt.Fprintf(w, "Post Method")
  default:
  	http.Error(w, "Method not supported", 405)
}}
