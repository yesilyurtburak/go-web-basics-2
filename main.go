package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = "8080"
const ipAddress = "127.0.0.1"

var url = fmt.Sprintf("%s:%s", ipAddress, portNumber)

// Serves home page by running this function whenever / is requested.
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

// Serves about page by running this function whenever /about is requested.
func AboutHandler(w http.ResponseWriter, r *http.Request) {

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
