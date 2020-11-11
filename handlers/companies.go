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


func Update(c *gin.Context) {
	var company model.Company
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
	} else {
		fmt.Printf("File %s uploaded successfully. \n", file.Filename)
	}

	
	for index, element := range Companies {
		//fmt.Println(index, element)
		if err := model.DB.Where("name = ? AND zip = ?", element.Name, element.Zip).First(&company).Error; err != nil {
			fmt.Println("Not found")
		}
		//fmt.Println(index, company)

		var input model.UpdateCompanyInput
		if err := c.ShouldBindJSON(&input); err != nil {
			fmt.Println(err)
		}	
		fmt.Println(index, company)
		model.DB.Model(&company).Updates(input)
	}


	c.JSON(http.StatusOK, gin.H{"data": Companies})

}
