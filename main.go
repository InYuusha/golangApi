package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
//book struct
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//init book var as slice
var books[] Book
//get books
func getBooks(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)

}
//get single book
func getBook(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r) //get params

	for _,item :=range books{
		if item.ID ==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
//create book
func createBook(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_=  json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books =append(books,book)
	json.NewEncoder(w).Encode(book)

}
//update book
func updateBook(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")

}
//delete book
func deleteBook(w http.ResponseWriter, r* http.Request){

	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range books {
		if item.ID==params["id"]{
			books = append(books[:index],books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main(){

	r:=mux.NewRouter()

	//mock data
	books = append(books,Book{ID:"1", Isbn:"44567", Title:"Book One", Author:&Author{Firstname:"John", Lastname:"Doe"}})

	books = append(books,Book{ID:"2", Isbn:"44678", Title:"Book Two", Author:&Author{Firstname:"Jk ", Lastname:"Rowlings"}})


	//route Handlers
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/book",createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}",deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))
}