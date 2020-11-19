package handlers

import (
	"net/http"
	"fmt"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

//Retorna todos os dados da tabela companies
func GetAll(c *gin.Context) {
	var companies []model.Company
	model.DB.Find(&companies)

	c.JSON(http.StatusOK, gin.H{"data": companies})
}

//Atualiza os dados da tabela companies que contém o Website vindo na requisição
func Update(c *gin.Context) {
	var company model.Company

	file, _ := c.FormFile("file")

	// Upload the file to specific dst.
	dst := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	//Lê o srquivo
	err := ReadCSV(file.Filename)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error on load file CSV")
		os.Exit(1)
	} else {
		fmt.Printf("File %s uploaded successfully. \n", file.Filename)
	}

	//Atualiza o banco
	for index, element := range Companies {

		fmt.Println(index, element)

		model.DB.Model(&company).Where("name = ? AND zip = ?", element.Name, element.Zip).Update("site", element.Site)

	}


	c.JSON(http.StatusOK, gin.H{"data": Companies})

}
