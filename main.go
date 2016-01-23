package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	
	http.HandleFunc("/", PageHandler)
	pageCol = myDB.Use("Pages")
	http.HandleFunc("/new", NewPageHandler)
	http.ListenAndServe(":3001", nil)

}

func NewPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprintf(w, "Post Method")
		docId, err := pageCol.Insert(map[string]interface{}{
			"hello": "world",
		})
		
		fmt.Println(docId)
		readBack, err := pageCol.Read(docId)

		if err != nil {
			panic(err)
		}

		fmt.Println(readBack)

		if err != nil {
			panic(err)
		}

	default:
		http.Error(w, "Method not supported", 405)
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case "GET":
			idKey := r.URL.Path[len("/"):]
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

