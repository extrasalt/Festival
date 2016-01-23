package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"log"
	"github.com/HouzuoGuo/tiedot/db"
	_ "github.com/HouzuoGuo/tiedot/dberr"
)

var pageCol *db.Col

func main() {
	myDB, err := db.OpenDB("data")

	if err != nil {
		panic(err)
	}

	if err := myDB.Create("Pages"); err != nil {
		fmt.Println(err)
	}
	
	pageCol = myDB.Use("Pages")

	http.HandleFunc("/p/", PageHandler)
	http.HandleFunc("/new", NewPageHandler)
	http.ListenAndServe(":3001", nil)

}

func NewPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Println("error parsing form")
		}
		docId, err := pageCol.Insert(map[string]interface{}{
			"title": r.PostFormValue("title"),
			"desc": r.PostFormValue("desc"),
			"date": r.PostFormValue("date"),
		})
		
		fmt.Println(docId)
		readBack, err := pageCol.Read(docId)

		if err != nil {
			panic(err)
		}

		fmt.Fprintln(w,readBack)

		if err != nil {
			panic(err)
		}

	case "GET":
		t, err := template.ParseFiles("templates/form.html")
		if err != nil {
			panic(err)
		}

		t.Execute(w, nil)

	default:
		http.Error(w, "Method not supported", 405)
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case "GET":
			idKey := r.URL.Path[len("/p/"):]
			fmt.Println(idKey)
			id, _ := strconv.Atoi(idKey)
			page, err := pageCol.Read(id)

			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, page)
		default:
			http.Error(w, "Methods not supported", 405)
	}
}

