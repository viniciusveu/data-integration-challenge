package handlers

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

func GetAll(c *gin.Context) {
	var companies []model.Company
	model.DB.Find(&companies)

	c.JSON(http.StatusOK, gin.H{"data": companies})
}

func GetOne(c *gin.Context) {
	
}

func Create(c *gin.Context) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	dst := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	err := ReadCSV(file.Filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error on load file CSV")
		os.Exit(1)
	}
	

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
}

func Delete(c *gin.Context) {

}