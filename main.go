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

	http.HandleFunc("/p/", PageHandler)
	http.HandleFunc("/new", NewPageHandler)
	http.ListenAndServe(":3001", nil)

}


