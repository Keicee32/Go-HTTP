package main

import (
	// "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
	"math/rand"
)
	
// Books Details with ID and all
type Books struct{
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Body	string	`json:"body"`

	// Name		string `json:"name"`
	// Author		string `json:"author"`
	// PublishedAt string `json:"published_at"`
}

var book []Books

func getPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func createPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var book Book _ = json.NewDecoder(r.Body).Decode(&book)
	book.Name = strconv.Itoa(rand.Intn(1000000))

	book = append(Books, book)

	json.NewEncoder(w).Encode(&book)
}

func getPost(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range book {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)

			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}

func updatePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range book {
		if item.ID == params["id"]{
			book = append(book[:index], book[index + 1:]...)

			var books Book _ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			book = append(book, books)

			json.NewEncoder(w).Encode(&books)
			return
		}
	}
	json.NewEncoder(w).Encode(book)
}

func deletePost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range book {
		if item.ID == params["id"]{
			book = append(book[:index], book[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(book)
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")


	http.ListenAndServe(":8000", router)
}