package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handle all request to root-url
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Woodcore Microservice")
}

//Articles will be returned with JSON files, first define an article structure
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type Articles []Article //array to fix in json article content

//demo up API endpoints that can help retrieve data
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Woodcore Test", Desc: "Test Description", Content: "Welcome to Woodcore"},
	}

	fmt.Println("Endpoint Hit: Microservices Endpoint")
	json.NewEncoder(w).Encode(articles)
}

//match url path
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET") //registering the new function - to tell the function all articles when it calls
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//entry point to rest server
func main() {
	handleRequests()
}
