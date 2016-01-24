package main

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	_ "github.com/HouzuoGuo/tiedot/dberr"
	"net/http"
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

	
	http.HandleFunc("/commit/", CommitHandler)
	http.HandleFunc("/e/", EditPageHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/gen", GeneratorHandler)
	http.HandleFunc("/p/", PageHandler)
	http.HandleFunc("/new", NewPageHandler)
	http.ListenAndServe(":3001", nil)

}
