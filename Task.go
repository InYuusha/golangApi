package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	CreatedOn time.Time `json:"createdon"`
}
var noteStore = make(map[string]Note)
var id int =0

//http POST Handler
func postNote(res http.ResponseWriter, req*http.Request){
	var note Note
	err:=json.NewDecoder(req.Body).Decode(&note)
	if err!=nil{
		http.Error(res,err.Error(),400)
	}
	note.CreatedOn=time.Now()
	id++
	k:=strconv.Itoa(id)
	noteStore[k]=note

	j,err:=json.Marshal(note)
	if err!=nil{
		panic(err)
	}
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(j)

}
//HTTP GET ALL

func getNotes(res http.ResponseWriter ,req *http.Request){
	var notes []Note
	for _, n:=range noteStore{
		notes=append(notes,n)
	}
	
	j,err:=json.Marshal(notes)
	if err!=nil{
		panic(err)
	}
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(j)
}

//HTTP PUT
func putNote(res http.ResponseWriter,req* http.Request){
	var newNote Note
	vars:=mux.Vars(req)
	id:=vars["id"]
	err:=json.NewDecoder(req.Body).Decode(&newNote)
	if err!=nil{
		panic(err)
	}
	if note,ok:=noteStore[id];ok{
		newNote.CreatedOn = note.CreatedOn
		delete(noteStore,id)
		noteStore[id]=newNote
	} else{
		log.Printf("Couldnt find the key %s to delete",id)
	}
	res.WriteHeader(http.StatusNoContent)
}
//HTTP DEL
func delNote(res http.ResponseWriter, req * http.Request){
	vars:=mux.Vars(req)
	id:=vars["id"]
	if _,ok:=noteStore[id];ok{
		delete(noteStore,id)
		res.WriteHeader(http.StatusGone)
	}else{
		http.Error(res,"Note Doesnt Exists",404)
	}
	
}

func main(){
	r:=mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes",getNotes).Methods("GET")
	r.HandleFunc("/api/notes",postNote).Methods("POST")
	r.HandleFunc("/api/notes/{id}",putNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}",delNote).Methods("DELETE")

	server:=&http.Server{
		Addr:":3000",
		Handler:r,
	}
	log.Println("Listening on port 3000")
	server.ListenAndServe()
}
