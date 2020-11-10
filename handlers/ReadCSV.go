package handlers

import(
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

var companies []model.Company
var company model.Company

func ReadCSV(csvName string) error  {
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

	if len(csvLines[1]) == 2 {
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
	} else {
		for index, line := range csvLines {
			if index == 0 {
				continue
			}
			// Código do merge aqui 
		}
	}

	

	jsonData, _ := json.Marshal(companies)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))

	return err
}
