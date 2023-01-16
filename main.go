package main

import (
	"errors"
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

func write(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg)) // converted from `string` to `[]byte` type... Because we can send the browser the data by using w.Write() which accepts `[]byte` type as an argument.
	checkError(err)
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return (x / y), nil
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	result, err := divide(5, 0)
	if err != nil {
		write(w, err.Error())
	} else {
		output := fmt.Sprintf("5 / 0 = %.2f\n", result)
		write(w, output) // send the formatted string to the browser by using function we created; write()
	}
}

func main() {
	// routing pages
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/divide", divideHandler)

	// listen to the port for incoming http requests.
	fmt.Printf("Listening traffic at %s\n", url)
	err := http.ListenAndServe(url, nil) // 2nd parameter is `nil` since we didn't send any information to the page.
	log.Fatal(err)
}
