package main

import (
	"fmt"
	"net/http"

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
