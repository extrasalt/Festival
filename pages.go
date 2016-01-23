package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Title string
	Desc  template.HTML
	Date  string
}

func NewPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Println("error parsing form")
		}
		docId, err := pageCol.Insert(map[string]interface{}{
			"title": r.PostFormValue("title"),
			"desc":  r.PostFormValue("desc"),
			"date":  r.PostFormValue("date"),
		})

		fmt.Println(docId)

		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/p/"+strconv.Itoa(docId), http.StatusFound)

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

func PageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idKey := r.URL.Path[len("/p/"):]
		fmt.Println(idKey)
		id, _ := strconv.Atoi(idKey)
		page, err := pageCol.Read(id)

		if err != nil {
			panic(err)
		}

		t, err := template.ParseFiles("templates/page.html")
		if err != nil {
			panic(err)
		}

		descMd := blackfriday.MarkdownCommon([]byte(page["desc"].(string)))
		t.Execute(w, &Page{
			Title: (page["title"]).(string),
			Desc:  template.HTML(descMd),
			Date:  (page["date"]).(string),
		})

	default:
		http.Error(w, "Methods not supported", 405)
	}
}
