package main

// importing shit
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Description"`
}

var Articles []Article

func allarticles(w http.ResponseWriter, e *http.Request) {
	articles := Articles{
		Article{Title: "First Article", Description: "blablablablabla", Content: "ehm"},
	}
	fmt.Println("allarticle end")
	json.NewEncoder(w).Encode(articles)
}

// The connection with the homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a homepage")
	fmt.Println("endpoint")
}

// This is the function that will handle the requests
func handlerequests() {
	http.HandleFunc("/", homepage)
	// Lets handle the articles func
	http.HandleFunc("/articles", allarticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// main function, nth special here
func main() {
	// calling the function
	handlerequests()
}
