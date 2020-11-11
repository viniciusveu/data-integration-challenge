package handlers

import(
	"fmt"
	"os"
	"encoding/csv"
	//"encoding/json"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

var Companies []model.UpdateCompanyInput

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

			company := model.Company{Name: line[0], Zip: line[1], Site: ""}
			model.DB.Create(&company)

		}
	} else {
		//var company model.Company

		for index, line := range csvLines {
			if index == 0 {
				continue
			}

			Companies = append(Companies, model.UpdateCompanyInput{Name: line[0], Zip: line[1], Site: line[2]})
		}
	}

	return err
}
