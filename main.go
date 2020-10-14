package main

// importing shit
import (
	"fmt"
	"log"
	"net/http"
)

// The connection with the homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a homepage")
	fmt.Println("endpoint")
}

// This is the function that will handle the requests
func handlerequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// main function, nth special here
func main() {
	// calling the function
	handlerequests()
}
