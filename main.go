package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func getAll(w http.ResponseWriter, r *http.Request) {

}

func getOne(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	  
    router.HandleFunc("/company", getAll).Methods("GET")
    router.HandleFunc("/company/{id}", getOne).Methods("GET")
    router.HandleFunc("/company", create).Methods("POST")
    router.HandleFunc("/company/{id}", delete).Methods("DELETE")

    
	fmt.Printf("Server running")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Rest API v1.0 - Mux Routers")

	handleRequests()
}