package main

import (
	"fmt"
	"os"
	model "github.com/viniciusveu/data-integration-challenge/models"
	handler "github.com/viniciusveu/data-integration-challenge/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	model.ConnectDatabase()

	err := handler.ReadCSV("./assets/q1_catalog.csv")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error on load file CSV")
		os.Exit(1)
	}

	router.GET("/company", handler.GetAll)
	router.PATCH("/company", handler.Update)

    
	fmt.Printf("Server running")
	router.Run(":8000")
}
