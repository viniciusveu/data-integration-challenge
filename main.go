package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"os"
	"encoding/csv"
	"encoding/json"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

var companies []model.Company
var company model.Company

func getAll(w http.ResponseWriter, r *http.Request) {

}

func getOne(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func readCSV(csvName string) error  {
	csvFile, err := os.Open(csvName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(csvLines)
	for index, line := range csvLines {
		if index == 0 {
			continue
		}
		// fmt.Println(len(line))
		model.DB.Create(&model.Company{Name: line[0], Zip: line[1], Site: ""})

		company.Name = line[0]
		company.Zip = line[1]
		company.Site = ""

		companies = append(companies, company)
	}

	jsonData, _ := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))

	return err
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

	model.ConnectDatabase()

	err := readCSV("./assets/q1_catalog.csv")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error on load file CSV")
		os.Exit(1)
	}

	handleRequests()
}