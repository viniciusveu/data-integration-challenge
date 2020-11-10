package main

import (
	"fmt"
	"os"
	model "github.com/viniciusveu/data-integration-challenge/models"
	handler "github.com/viniciusveu/data-integration-challenge/handlers"
	"github.com/gin-gonic/gin"
)

func handleRequests() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	model.ConnectDatabase()

	router.GET("/company", handler.GetAll)
	router.GET("/company/:id", handler.GetOne)
	router.POST("/company", handler.Create)
	router.DELETE("/company/:id", handler.Delete)
    
	fmt.Printf("Server running")
	router.Run(":8080")
}

func main() {
	fmt.Println("Rest API v1.0 - Mux Routers")

	handleRequests()

	err := handler.ReadCSV("./assets/q1_catalog.csv")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error on load file CSV")
		os.Exit(1)
	}

}