package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"strconv"
	"os"
)

func CommitHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idKey := r.URL.Path[len("/commit/"):]
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
		
		file, err := os.Create("public/words.html")
		
		if err != nil {
			panic(err)
		}

		descMd := blackfriday.MarkdownCommon([]byte(page["desc"].(string)))
		t.Execute(file, &Page{
			Title: (page["title"]).(string),
			Desc:  template.HTML(descMd),
			Date:  (page["date"]).(string),
		})
		
		fmt.Fprintln(w, "Written to file")
	default:
		http.Error(w, "Methods not supported", 405)
	}
}

