package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func GeneratorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("templates/generator.html")

		if err != nil {
			log.Println("Failed to parse files")
		}
		t.Execute(w, nil)

	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Println("Failed to get post values")
		}

		rawText := r.PostFormValue("desc")

		dateField := ParseDate(rawText)
		//desc := blackfriday.MarkdownCommon([]byte(rawText))

		docId, err := pageCol.Insert(map[string]interface{}{
			"title": "From Generator",
			"desc":  rawText,
			"date":  dateField,
		})

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/p/"+strconv.Itoa(docId), http.StatusFound)
	default:
		http.Error(w, "Method not allowed", 405)

	}
}

func ParseDate(sample string) time.Time {
	datePattern, err := regexp.Compile(`(\d{1,2}\b\D{3,9}\b\d{4})|([a-zA-Z]{3,9}\s\d{1,2}\s\d{4})`)

	if err != nil {
		panic(err)
	}

	colloquialPattern, err := regexp.Compile(`(\d{1,2}\b\D{3,9}\b\d{4})`)

	if err != nil {
		panic(err)
	}

	americanPattern, err := regexp.Compile(`([a-zA-Z]{3,9}\s\d{1,2}\s\d{4})`)
	if err != nil {
		panic(err)
	}
	var t time.Time
	dateString := datePattern.FindString(sample)
	switch {
	case americanPattern.MatchString(dateString):
		t, _ = time.Parse("January 2 2006", datePattern.FindString(sample))

	case colloquialPattern.MatchString(dateString):
		t, _ = time.Parse("2 January 2006", datePattern.FindString(sample))
	}
	return t
}
