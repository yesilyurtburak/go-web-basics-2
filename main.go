package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const portNumber = "8080"
const ipAddress = "127.0.0.1"

var url = fmt.Sprintf("%s:%s", ipAddress, portNumber)

// This function aims to reduce repetitive code.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmpl),
		"./templates/base.layout.gotmpl",
	}
	// this creates a new template from the given files.
	parsedTemplate, err := template.ParseFiles(templates...)
	checkError(err)
	// write the output of parsed template to the response writer `w`.
	// in other words; send the template data to the browser by using w.
	err = parsedTemplate.Execute(w, nil)
	checkError(err)
}

// Serves home page by running this function whenever / is requested.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gotmpl")
}

// Serves about page by running this function whenever /about is requested.
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gotmpl")
}

func main() {
	// routing pages
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	// listen to the port for incoming http requests.
	fmt.Printf("Listening traffic at %s\n", url)
	err := http.ListenAndServe(url, nil) // 2nd parameter is `nil` since we didn't send any information to the page.
	log.Fatal(err)
}
