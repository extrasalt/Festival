package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"crypto/rand"
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
		filename := randStr()+".html"
		file, err := os.Create("public/"+filename)
		defer file.Close()

		if err != nil {
			panic(err)
		}

		descMd := blackfriday.MarkdownCommon([]byte(page["desc"].(string)))
		t.Execute(file, &Page{
			Title: (page["title"]).(string),
			Desc:  template.HTML(descMd),
			Date:  (page["date"]).(string),
		})

		http.Redirect(w, r, "/"+filename, http.StatusFound)
		fmt.Fprintln(w, "Written to file")
	default:
		http.Error(w, "Methods not supported", 405)
	}
}

func randStr() string {

	dictionary := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, 8)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
